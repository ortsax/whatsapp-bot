---
layout: page
title: Configuration
nav_order: 3
---

# Configuration
{: .no_toc }

<details open markdown="block">
  <summary>Table of contents</summary>
  {: .text-delta }
- TOC
{:toc}
</details>

---

## Environment variables

Alphonse reads a `.env` file in the working directory (falling back to `.env.example`). You can also export variables directly in your shell or service file.

| Variable | Default | Description |
|---|---|---|
| `DATABASE_URL` | `database.db` | SQLite file path **or** a `postgres://…` connection string |

### SQLite (default)

Any value that does not start with `postgres://` or `postgresql://` is treated as a SQLite file path:

```env
DATABASE_URL=database.db
```

Alphonse appends WAL mode, busy-timeout, and other recommended pragmas automatically — you do not need to configure them manually.

### PostgreSQL

```env
DATABASE_URL=postgres://user:password@localhost:5432/alphonse?sslmode=disable
```

Install the `libpq` client library on the host if it is not already present.

---

## Bot settings

Bot settings are persisted per-owner in the database (table `bot_settings`) and can be changed at runtime through chat commands without restarting the bot.

### Prefixes

The character(s) that must precede a command name for Alphonse to recognise it.

```
.setprefix .
```

Multiple prefixes are separated by spaces. Use the token `empty` to allow commands with no prefix at all:

```
.setprefix . ! empty
```

Default: `.`

### Mode

Controls who can invoke commands.

| Value | Who can use commands |
|---|---|
| `public` | Everyone (default) |
| `private` | Sudo users only |

```
.setmode private
.setmode public
```

### Sudo users

Sudo users bypass permission checks and have access to all commands regardless of mode.

```
.setsudo 2348012345678      # add a sudo user
.delsudo 2348012345678      # remove a sudo user
.getsudo                    # list all sudo users
```

The bot owner (the phone number used to pair) is automatically added as a sudo user on first startup.

### Language

Alphonse ships with 10 built-in languages. Change the language for all bot responses:

```
.lang en    # English (default)
.lang es    # Spanish
.lang pt    # Portuguese
.lang ar    # Arabic
.lang hi    # Hindi
.lang fr    # French
.lang de    # German
.lang ru    # Russian
.lang tr    # Turkish
.lang sw    # Swahili
```

### Group chat responses

Disable or re-enable Alphonse responding to commands in group chats:

```
.shh        # disable group responses
.shh off    # re-enable group responses
```

### Anti-delete

When enabled, Alphonse forwards deleted messages to the bot owner:

```
.antidelete on
.antidelete off
```

### Banned users

Banned users are silently ignored — no response is sent.

```
.ban 2348012345678
.unban 2348012345678
```

### Disabling individual commands

```
.disable ping
.enable ping
```

---

## Build-time flags

When building from source you can inject the source directory path so that `alphonse --update` knows where to find the repo:

```bash
go build \
  -ldflags="-s -w -X main.sourceDir=/opt/alphonse/src" \
  -trimpath \
  -o alphonse \
  .
```

The install scripts do this automatically.
