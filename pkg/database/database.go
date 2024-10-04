package database

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"runtime"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
	osPath := setPath()

	dbName, nameErr := makeDBName()
	if nameErr != nil {
		return nil, nameErr
	}

	dbPath := osPath + dbName

	db, dbErr := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if dbErr != nil {
		return nil, dbErr
	}

	return db, nil
}

func setPath() string {
	localPath := ""
	if runtime.GOOS == "windows" {
		localPath = `%LOCALAPPDATA%\passh`
	} else {
		localPath = `~/.local/share/passh`
	}
	return localPath
}

func makeDBName() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot set home directory: %w", err)
	}

	hash := sha1.New()
	hash.Write([]byte(homeDir))
	sha1Hash := hex.EncodeToString(hash.Sum(nil))

	return sha1Hash, nil
}
