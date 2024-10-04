# passh

![Passh logo](./assets/logo-no-background.png)

**Pass** is an open-source password manager written in Go. It uses the `pass` command line tool to easily store and retreive your passwords.

## Installation

```bash
go install github.com/mclacore/passh@latest
```

## Usage

### Password
Generate a new password:
```bash
passh pass new
```

### Login
Create a new login (with specific password):
```bash
passh login new --item-name Stackoverflow --username mclacore --password 123456 --url https://stackoverflow.com
```

Create a new login (with generated password):
```bash
passh login new --item-name Stackoverflow --username mclacore --url https://stackoverflow.com
```

Create a new login with no password:
```bash
passh login new --item-name Stackoverflow --username mclacore --url https://stackoverflow.com --no-password
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

- Need to decrypt on db call, then encrypt after db call
- Need fast decrypt/encrypt
- Double encryption on passwords?
- Iterate through itemNames when doing a GET and list all matching items
- Add collection argument to login
- Create a default collection for initial login items
- Move login items to another collection

Move login item to another collection
```bash
passh login move --from-collection Personal --to-collection Work
```
-->
