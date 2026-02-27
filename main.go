package main

import (
	"context"
	"flag"
	"fmt"
	"orstax/plugins"
	"orstax/store"
	"orstax/store/sqlstore"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"go.mau.fi/whatsmeow"
	waLog "go.mau.fi/whatsmeow/util/log"

	_ "github.com/lib/pq"
	_ "modernc.org/sqlite"
)

// loadEnv loads .env if present, otherwise falls back to .env.example.
func loadEnv() {
	if err := godotenv.Load(".env"); err != nil {
		_ = godotenv.Load(".env.example")
	}
}

// dbConfig returns the sql dialect and connection address derived from DATABASE_URL.
// A bare filename (no scheme) or a path ending in .db is treated as SQLite;
// anything starting with postgres:// or postgresql:// is treated as PostgreSQL.
func dbConfig() (dialect, addr string) {
	url := os.Getenv("DATABASE_URL")
	if url == "" {
		url = "database.db"
	}

	if strings.HasPrefix(url, "postgres://") || strings.HasPrefix(url, "postgresql://") {
		return "postgres", url
	}

	// SQLite – build the connection string with recommended pragmas.
	// Strip a leading "file:" if present so we can normalise the path.
	path := strings.TrimPrefix(url, "file:")
	addr = "file:" + path +
		"?_pragma=foreign_keys(1)" +
		"&_pragma=journal_mode(WAL)" +
		"&_pragma=synchronous(NORMAL)" +
		"&_pragma=busy_timeout(10000)" +
		"&_pragma=cache_size(-64000)" +
		"&_pragma=mmap_size(2147483648)" +
		"&_pragma=temp_store(MEMORY)"
	return "sqlite", addr
}

// getDevice returns the device for the given phone number.
// If phone is empty it falls back to the first stored device (or a new one).
// If phone is provided and no matching device exists, a new (unpaired) device is returned.
func getDevice(ctx context.Context, container *sqlstore.Container, phone string) (*store.Device, error) {
	if phone == "" {
		return container.GetFirstDevice(ctx)
	}

	devices, err := container.GetAllDevices(ctx)
	if err != nil {
		return nil, err
	}
	for _, dev := range devices {
		if dev.ID == nil {
			continue
		}
		// Device JID User field may be "phone.deviceIndex" – compare only the phone part.
		userPhone := strings.SplitN(dev.ID.User, ".", 2)[0]
		if userPhone == phone {
			return dev, nil
		}
	}
	// No existing session for this number – return a fresh device for pairing.
	return container.NewDevice(), nil
}

func main() {
	loadEnv()

	phoneArg := flag.String("phone-number", "", "Phone number (international format) used to identify or pair a device")
	flag.Parse()

	dbLog := waLog.Stdout("Database", "ERROR", true)
	ctx := context.Background()

	dialect, dbAddr := dbConfig()

	container, err := sqlstore.New(ctx, dialect, dbAddr, dbLog)
	if err != nil {
		panic(err)
	}

	// Create the bot_settings table (no user needed yet).
	if err := plugins.InitDB(container.DB()); err != nil {
		panic(fmt.Errorf("settings db init: %w", err))
	}

	// Wire LID resolver (owner phone resolved after Connect).
	plugins.InitLIDStore(container.LIDMap, "")

	deviceStore, err := getDevice(ctx, container, *phoneArg)
	if err != nil {
		panic(err)
	}

	clientLog := waLog.Stdout("Client", "ERROR", true)
	client := whatsmeow.NewClient(deviceStore, clientLog)
	client.AddEventHandler(plugins.NewHandler(client))

	err = client.Connect()
	if err != nil {
		panic(err)
	}

	if client.Store.ID == nil {
		if *phoneArg == "" {
			fmt.Println("No session found. Please provide a phone number using --phone-number")
			return
		}

		fmt.Println("Waiting 10 seconds before generating pairing code...")
		time.Sleep(10 * time.Second)

		code, err := client.PairPhone(ctx, *phoneArg, true, whatsmeow.PairClientChrome, "Chrome (Linux)")
		if err != nil {
			panic(err)
		}
		fmt.Printf("Your pairing code is: %s\n", code)
	} else {
		ownerPhone := strings.SplitN(client.Store.ID.User, ".", 2)[0]
		plugins.InitLIDStore(container.LIDMap, ownerPhone)
		if err := plugins.InitSettings(ownerPhone); err != nil {
			panic(fmt.Errorf("settings load: %w", err))
		}
		plugins.BootstrapOwnerSudoers()
		fmt.Println("Already logged in.")
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	client.Disconnect()
}
