# kk - The CLI Fuzzy-Finder Scratchpad

[![Go Report Card](https://goreportcard.com/badge/github.com/moosemanf/kk)](https://goreportcard.com/report/github.com/moosemanf/kk)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

`kk` is a tiny, lightning-fast command-line tool for storing and retrieving the everyday snippets, notes, and commands you need but always forget. Stop searching and start finding.

![kk-demo](https://user-images.githubusercontent.com/your-image-host/kk-demo.gif)
*(A demo GIF showing `kk set "staging-url"` and then `kk pick` would be perfect here)*

## Why `kk`?

Tired of searching through your shell history, Slack messages, or a heavy notes app just to find...
* ...the URL for a feature branch you deployed?
* ...the Jira ticket number for your current epic?
* ...that one-liner `ffmpeg` command you use once a month?
* ...the cost center number you need for an invoice?
* ...an email address or a specific file path?

`kk` is your personal, encrypted scratchpad. It uses `fzf` to give you instant access to your information right when you need it, all without leaving your terminal.

## Features

* **Instant Fuzzy-Finding**: Uses `fzf` to instantly search and select the note you need.
* **Clipboard-Ready**: Instantly copy any snippet to your clipboard with the `-c` flag.
* **Encrypted On-Disk**: Your notes are encrypted using modern [`age`](https://github.com/FiloSottile/age) encryption. It's not a password manager, but your data has a strong layer of privacy on your disk.
* **Simple & Portable**: Your entire collection of notes is in a single, portable file.

## Prerequisites

Before using `kk`, you need to install two external tools it depends on:

1.  **`age`**: To generate the cryptographic keys for encrypting your notes file.
    * Installation instructions: [github.com/FiloSottile/age#installation](https://github.com/FiloSottile/age#installation)

2.  **`fzf`**: For the interactive fuzzy-finding menu.
    * Installation instructions: [github.com/junegunn/fzf#installation](https://github.com/junegunn/fzf#installation)

## Installation

Assuming you have a working Go environment:
```sh
go install [github.com/moosemanf/kk@latest](https://github.com/moosemanf/kk@latest)
```
## Setup and Prerequisites

Before you can use `kk`, you need to install its dependencies and generate a cryptographic key.

### 1. Install Dependencies

You must install two command-line tools:

* **`age`**: Used to generate keys and handle the underlying file encryption.
    * **Why?**: `kk` uses the `age` library to ensure your notes are stored securely and privately on your disk.
    * **Installation**: Follow the official instructions at [github.com/FiloSottile/age#installation](https://github.com/FiloSottile/age#installation).

* **`fzf`**: The interactive fuzzy-finder used by the `pick` command.
    * **Why?**: `fzf` provides the lightning-fast menu that lets you search for your notes without remembering the exact key.
    * **Installation**: Follow the official instructions at [github.com/junegunn/fzf#installation](https://github.com/junegunn/fzf#installation).

### 2. Generate Your Key and Recipient Files

You only need to do this setup once. `kk` looks for its keys in specific locations in your home directory.

**Step 2a: Create the key file**

Use the `age-keygen` command to generate your primary key. This key is your identity and is used to decrypt your notes.

```sh
# First, create the directory where age keys are typically stored
mkdir -p ~/.age

# Now, generate your key file inside that directory
age-keygen -o ~/.age/key.txt
```

This command creates a file at `~/.age/key.txt`. It will also print the corresponding public key to your terminal. The public key is not secret and starts with age1....

**Step 2b: Create the recipient file**

Next, you must tell kk who to encrypt notes for. You do this by creating a recipient file that contains your public key.

Copy the public key (the line starting with age1...) from the output of the previous command and save it to a new file at `~/.age/recipient.txt`.

```sh
# Replace with the actual public key you just generated
echo "age1ql3z7hjy54pw3hyww5mmdo5s5j9yr6hfg2cln5vscf2jshj5dkzspchv4d" > ~/.age/recipient.txt
```

**Setup is complete!** Your kk tool is now configured and ready to use. It relies on these files:

`~/.age/key.txt`	Your private key. Keep this file safe, just like an SSH key. Used for decryption.
`~/.age/recipient.txt`	Your public key. Used for encrypting new notes.
`~/.kk.age`	Your encrypted notes file. This is created and managed automatically.

## Usage
Here is how to interact with your kk notes database.

### Adding a Snippet (`kk set <key>`)
Use `kk set <key>` to store a new piece of information. The command will then prompt you to enter the value. This method is secure because the value you enter won't be saved in your shell's history.

Example:
```sh
$ kk set main-project-epic
Value: PROJ-1234
```
Your snippet is now saved and encrypted.

### Retrieving a Snippet Directly (`kk get <key>`)
If you know the exact key for a snippet, you can retrieve it directly using `kk get <key>`.

Example:
```sh
$ kk get main-project-epic
PROJ-1234
```

### Finding a Snippet Interactively (`kk pick`)
This is the most powerful command for retrieval. Use `kk pick` when you want to search for a snippet but don't remember its exact key. It opens an interactive fuzzy-finder menu.

Example:
```sh
$ kk pick
```
Once you select an item from the menu, its value will be printed to the console.

### Copying to Clipboard

To copy the value of a selected snippet directly to your clipboard, use the `-c` or `--clip` flag.

Example:
```sh
$ kk pick -c
```

A confirmation message will be printed, and the value will be ready to paste elsewhere.