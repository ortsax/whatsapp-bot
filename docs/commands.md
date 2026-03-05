---
layout: page
title: Command Reference
nav_order: 4
---

# Command Reference
{: .no_toc }

All commands use the configured prefix (default `.`). Commands marked **sudo** require the sender to be a sudo user or the bot owner. Commands marked **admin** require both the bot and the sender to be group administrators.

<details open markdown="block">
  <summary>Table of contents</summary>
  {: .text-delta }
- TOC
{:toc}
</details>

---

## Utility

### menu / help

Show the interactive command menu grouped by category.

```
.menu
.help
```

### ping

Respond with `pong` and the round-trip latency.

```
.ping
```

### lang

Change the language used for all bot responses.

```
.lang <code>
```

Supported codes: `en`, `es`, `pt`, `ar`, `hi`, `fr`, `de`, `ru`, `tr`, `sw`.

### update {#update}

Pull the latest source from GitHub and rebuild the binary in-place. **Sudo only.**

```
.update
```

---

## Settings (sudo only)

### setprefix

Change the command prefix(es). Separate multiple prefixes with spaces. Use `empty` to allow prefix-less commands.

```
.setprefix .
.setprefix . ! empty
```

### setmode / getmode

Switch between `public` (everyone) and `private` (sudo users only) modes.

```
.setmode private
.setmode public
```

### setsudo / delsudo / getsudo

Manage the sudo user list.

```
.setsudo <phone>     # grant sudo
.delsudo <phone>     # revoke sudo
.getsudo             # list sudo users
```

### ban / unban

Silently ignore (or unignore) a user.

```
.ban <phone>
.unban <phone>
```

### disable / enable

Disable or re-enable a specific command by name.

```
.disable <command>
.enable <command>
```

---

## Moderation

### antidelete

When on, forwards deleted messages to the bot owner.

```
.antidelete on
.antidelete off
```

### antilink

Remove messages containing WhatsApp group invite links and optionally warn/kick the sender. **Admin, group only.**

```
.antilink on
.antilink off
```

### antispam

Automatically mute members who send messages too rapidly. **Admin, group only.**

```
.antispam on
.antispam off
```

### anticall

Reject or ignore incoming calls automatically.

```
.anticall on
.anticall off
.anticall reject     # reject calls with a decline response
.anticall block      # block callers automatically
```

### antistatus

Block automatic status-update receipts from being sent.

```
.antistatus on
.antistatus off
```

### antiword

Remove messages containing specified words and warn the sender. **Admin, group only.**

```
.antiword add <word>
.antiword remove <word>
.antiword list
.antiword on
.antiword off
```

### antivv

Auto-open (view-once bypass) and forward view-once messages to the sender.

```
.antivv on
.antivv off
```

Manually open a specific view-once message (reply to it):

```
.vv
```

### warn

Issue a warning to a user (reply to their message). Three warnings trigger an automatic kick. **Admin, group only.**

```
.warn
```

### report

Report a message to the bot owner (reply to the message).

```
.report
```

---

## Group Administration (admin, group only)

### promote / demote

Change a member's admin status.

```
.promote @user
.demote @user
```

### kick

Remove a member from the group.

```
.kick @user
```

### kickall

Remove all non-admin members from the group. **Sudo only.**

```
.kickall
```

### mute / unmute

Restrict all members from sending messages (mute the group).

```
.mute
.mute off
```

### shh

Toggle whether Alphonse responds to commands in group chats.

```
.shh          # disable group responses
.shh off      # re-enable
```

### newgc

Create a new WhatsApp group with participants taken from a reply or mentions.

```
.newgc <group name>
```

### filter / gfilter / dfilter

Auto-reply to matching keywords.

```
.filter <trigger> | <response>     # add a personal filter
.gfilter <trigger> | <response>    # add a group-wide filter
.dfilter <trigger>                 # delete a filter
```

---

## Chat utilities

### del

Delete a message (reply to it). The bot must be admin in groups.

```
.del
```

### star / unstar

Star or unstar a message (reply to it).

```
.star
.unstar
```

### pin / unpin

Pin or unpin a message (reply to it). **Admin, group only.**

```
.pin
.unpin
```

### archive

Archive a chat.

```
.archive
```

### block / unblock

Block or unblock a contact.

```
.block
.unblock
```

### afk

Mark yourself as away. Alphonse will auto-reply to messages while you are AFK.

```
.afk                    # go AFK
.afk <reason>           # go AFK with a reason
.afk off                # return from AFK
```

---

## Media

### mp3

Extract audio from a video message as an MP3 file (reply to the video).

```
.mp3
```

### trim

Trim a video to a time range (reply to the video).

```
.trim <start> <end>
```

Example: `.trim 0:05 0:30`

### black

Remove black borders from a video (reply to the video).

```
.black
```

---

## Status

### autosavestatus

Automatically save all contact status updates to your phone.

```
.autosavestatus on
.autosavestatus off
```

### autolikestatus

Automatically react with a ❤️ to every contact status update.

```
.autolikestatus on
.autolikestatus off
```

---

## AI

### meta

Send a message to Meta AI and receive a reply in the chat.

```
.meta <prompt>
```

Example: `.meta What is the capital of France?`
