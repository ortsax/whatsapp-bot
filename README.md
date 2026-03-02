# whatsapp-bot

A WhatsApp bot built in Go using [whatsmeow](https://github.com/tulir/whatsmeow). Connects via phone-number pairing (no QR scan), persists sessions in SQLite or PostgreSQL, and supports an extensible command plugin system.

## Installation

Pick the script for your platform and run it with elevated privileges. The script handles everything — Go, Git, cloning, building, and PATH setup.

**Windows** (PowerShell as Administrator)

```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force
irm https://raw.githubusercontent.com/ortsax/whatsapp-bot/master/scripts/install.ps1 | iex
```

**Linux**

```bash
sudo bash <(curl -fsSL https://raw.githubusercontent.com/ortsax/whatsapp-bot/master/scripts/install-linux.sh)
```

**macOS**

```bash
sudo bash <(curl -fsSL https://raw.githubusercontent.com/ortsax/whatsapp-bot/master/scripts/install-mac.sh)
```

Once complete you will see:

```
  Ortsax is installed!

  Run with      ortsax --phone-number <number>
  Update with   ortsax --update
```

> Open a new terminal after install so the updated PATH takes effect.

## First run — pairing

```bash
ortsax --phone-number <international-format-number>
```

A pairing code will be printed. On your phone go to **WhatsApp → Linked Devices → Link a Device → Link with phone number instead** and enter the code.

## Subsequent runs

```bash
ortsax
```

Press `Ctrl+C` to disconnect.

## Database

By default Ortsax uses a local SQLite file. To use PostgreSQL, create a `.env` file next to the binary:

```env
# SQLite (default)
DATABASE_URL=database.db

# PostgreSQL
DATABASE_URL=postgres://user:pass@localhost:5432/mydb
```

Pulls the latest source and rebuilds the binary in-place. Stop the bot first on Windows before updating.

## License

See [LICENSE](LICENSE).
