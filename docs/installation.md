---
layout: page
title: Installation
nav_order: 2
---

# Installation
{: .no_toc }

<details open markdown="block">
  <summary>Table of contents</summary>
  {: .text-delta }
- TOC
{:toc}
</details>

---

## Prerequisites

Before installing Alphonse you need:

- **Git** — [git-scm.com](https://git-scm.com) (Windows) or your distro's package manager
- **Go 1.21+** — [go.dev/dl](https://go.dev/dl) (the Linux and macOS scripts install Go automatically if missing)
- A WhatsApp account with a valid phone number

---

## One-liner install scripts

The easiest way to install Alphonse is via the platform-specific install script. Each script:

1. Checks for (and optionally installs) Go and Git
2. Clones the repository to a system directory
3. Builds the `alphonse` binary with optimised flags
4. Adds the install directory to your system `PATH`

### Windows

Open **PowerShell as Administrator** and run:

```powershell
Set-ExecutionPolicy Bypass -Scope Process -Force
irm https://raw.githubusercontent.com/ortsax/Alphonse/master/scripts/install.ps1 | iex
```

The binary is installed to `%ProgramData%\alphonse\alphonse.exe` and the directory is added to the machine-wide `PATH`.

> Restart your terminal after installation for the PATH change to take effect.

### Linux

```bash
sudo bash <(curl -fsSL https://raw.githubusercontent.com/ortsax/Alphonse/master/scripts/install-linux.sh)
```

The binary is installed to `/usr/local/bin/alphonse`. Source code lives in `/opt/alphonse/src`.

> If Go was installed by the script, run `source /etc/profile.d/go.sh` or open a new shell.

### macOS

```bash
sudo bash <(curl -fsSL https://raw.githubusercontent.com/ortsax/Alphonse/master/scripts/install-mac.sh)
```

Go is installed via Homebrew if available, otherwise downloaded directly from go.dev. The binary lands in `/usr/local/bin/alphonse`.

---

## Manual build (from source)

```bash
git clone https://github.com/ortsax/Alphonse.git
cd Alphonse
go build -o alphonse .
```

For a production build matching what the install scripts produce:

```bash
go build \
  -ldflags="-s -w -X main.sourceDir=$(pwd)" \
  -trimpath \
  -o alphonse \
  .
```

---

## Pairing your phone

Once installed, pair a WhatsApp account:

```
alphonse --phone-number <number>
```

`<number>` must be in **international format** without the leading `+` — for example `2348012345678` for a Nigerian number.

Alphonse will print a pairing code:

```
Your pairing code is: ABCD-1234
```

On your phone open **WhatsApp → Settings → Linked Devices → Link a Device → Link with phone number instead** and enter the code. The session is saved to the database automatically; subsequent runs of `alphonse` will reconnect without a code.

---

## Session management

| Command | Effect |
|---|---|
| `alphonse --list-sessions` | List all paired sessions |
| `alphonse --delete-session <number>` | Permanently remove a session |
| `alphonse --reset-session <number>` | Clear a session so it can be re-paired |

---

## Updating

```
alphonse --update
```

This pulls the latest source from GitHub, rebuilds the binary in-place, and exits. Restart `alphonse` to run the new version. The same operation is available as a chat command — see [Command Reference › update](commands#update).

---

## Database

By default Alphonse stores everything in `database.db` (SQLite) in the working directory. To use PostgreSQL set the `DATABASE_URL` environment variable:

```bash
DATABASE_URL=postgres://user:pass@localhost/alphonse alphonse
```

See [Configuration](configuration#database) for details.
