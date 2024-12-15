# passh

![Passh logo](./assets/logo-no-background.png)

**Passh** is an open-source password manager written in Go. It uses the `passh` command line tool to easily store and retrieve your passwords.

## Installation

```bash
go install github.com/mclacore/passh@latest
```

## Usage

### Password

Create secure passwords with customizable requirements such as length, character types, and complexity.

Generate a new password:

```bash
passh pass new
```

Specify password requirements:

```bash
passh pass new --length 40 \
    --uppercase true \
    --exclude-lowercase false \
    --numbers true \
    --special true
```

### Login

Store and manage login credentials, including usernames, passwords, and URLs, with options for generating or using specific passwords.

Create a new login (with specific password):

```bash
passh login new --item-name Stackoverflow \
    --username mclacore \
    --password 123456 \
    --url https://stackoverflow.com
```

Create a new login (with generated password):

```bash
passh login new --item-name Stackoverflow \
    --username mclacore \
    --url https://stackoverflow.com
```

Create a new login with no password:

```bash
passh login new --item-name Stackoverflow \
    --username mclacore \
    --url https://stackoverflow.com \
    --no-password
```

List all logins:

```bash
passh login list
```

Get login details (password hidden):

```bash
passh login get --item-name Stackoverflow
```

Get login details (password shown):

```bash
passh login get --item-name Stackoverflow --show-password
```

Update a login:

```bash
passh login update --item-name Stackoverflow --password 654321
```

Delete a login:

```bash
passh login delete --item-name Stackoverflow
```

### Collection

Organize login items into collections, grouping them by categories like work, personal, or school for better management.

Create new collection

```bash
passh collection new --collection-name Work
```

List collections

```bash
passh collection list
```

Delete collection

```bash
passh collection delete --collection-name Work
```
<!---
TODO:

- Create a way to look for passh files
- Need to decrypt on db call, then encrypt after db call
- Need fast decrypt/encrypt
- Double encryption on passwords?
- Set up SSH server
- Set up docker image

-->
