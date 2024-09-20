# passh

![Passh logo](./assets/logo-no-background.png)

**Pass** is an open-source password manager written in Go. It uses the `pass` command line tool to easily store and retreive your passwords.

## Installation

```bash
go install github.com/mclacore/passh@latest
```

## Usage

Generate a new password:
```bash
passh pass new
```

Create a new login (with specific password):
```bash
passh login new --item-name Stackoverflow --username mclacore --password 123456 --url https://stackoverflow.com
```

Create a new login (with generated password):
```bash
passh login new --item-name Stackoverflow --username mclacore --url https://stackoverflow.com
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
passh login rm --item-name Stackoverflow
```
