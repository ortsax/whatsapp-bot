---
layout: home
title: Home
nav_order: 1
---

# Alphonse

A self-hosted WhatsApp bot written in Go, built on [whatsmeow](https://github.com/tulir/whatsmeow). Alphonse connects via **phone-number pairing** — no QR code, no phone scan — and runs as a single statically-linked binary on Windows, Linux, and macOS.

## What Alphonse can do

- **Group moderation** — anti-link, anti-spam, anti-delete, anti-call, warn system, kick/mute/promote/demote members
- **Media processing** — extract audio as MP3, trim video clips, remove black borders from videos
- **Status automation** — automatically save and/or like contact status updates
- **AI assistant** — chat with Meta AI directly from WhatsApp via the `meta` command
- **Flexible settings** — per-owner configuration stored in SQLite or PostgreSQL; change prefixes, language, mode, and sudo users without restarting
- **Self-update** — run `alphonse --update` (or `.update` in chat) to pull the latest source and rebuild in-place
- **10 languages** — English, Spanish, Portuguese, Arabic, Hindi, French, German, Russian, Turkish, Swahili

## Getting started

Head to the [Installation](installation) page for a step-by-step guide, then come back here for [Configuration](configuration) and the full [Command Reference](commands).

## Quick links

| Page | What's in it |
|---|---|
| [Installation](installation) | One-liner scripts for Windows, Linux, macOS |
| [Configuration](configuration) | Environment variables, database options, settings commands |
| [Command Reference](commands) | Every built-in command with usage examples |
| [Plugin Development](plugins) | Writing and registering your own commands |

## Requirements

- Go 1.21 or later (the install scripts handle this automatically)
- Git
- A WhatsApp account with an active phone number

## License

Alphonse is released under the [MIT License](https://github.com/ortsax/Alphonse/blob/master/LICENSE).
