# whatsapp-bot

A WhatsApp bot built in Go using [whatsmeow](https://github.com/tulir/whatsmeow). Connects via phone-number pairing (no QR scan), persists sessions in SQLite, and supports an extensible command plugin system.

## Requirements

- Go 1.25+

## Setup

```bash
git clone https://github.com/your-username/whatsapp-bot.git
cd whatsapp-bot
go build -o orstax .
```

### Pair a new device

```bash
./orstax --phone-number <international-format-number>
```

A pairing code will be printed. On your phone go to **WhatsApp → Linked Devices → Link a Device → Link with phone number instead** and enter the code.

### Run (already paired)

```bash
./orstax
```

Press `Ctrl+C` to disconnect.
