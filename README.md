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
passh login del --item-name Stackoverflow
```

<!---
Setting up collections:

- Collections will be their own tables
- Database file will need to be stored _somewhere_
- Need to decrypt on db call, then encrypt after db call
- Need fast decrypt/encrypt
- Double encryption on passwords?

# Create new collection
```bash
passh collection (col) new --name/n colName
```

# List collections
```bash
passh collection list
```

# Get login items in collection
```bash
passh login list --collection/c colName
```

# Migrate item from 1 collection to another
# Perfoms:
1. Get login item
2. Create login item in newColName
3. Verify login item in newColName
4. Delete login item in oldColName
```bash
passh login migrate --from-collection/f oldColName --to-collection/t newColName
```
-->
