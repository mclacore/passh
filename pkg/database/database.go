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
	osPath, pathErr := setPath()
	if pathErr != nil {
		return nil, pathErr
	}

	dbName, nameErr := makeDBName()
	if nameErr != nil {
		return nil, nameErr
	}

	dbPath := osPath + dbName

	if _, err := os.Stat(dbPath); err != nil {
		dbFile, fileErr := os.Create(dbPath)
		if fileErr != nil {
			return nil, fmt.Errorf("cannot create database: %w", fileErr)
		}
		defer dbFile.Close()
	}

	db, dbErr := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if dbErr != nil {
		return nil, dbErr
	}

	return db, nil
}

func setPath() (string, error) {
	localPath := ""
	if runtime.GOOS == "windows" {
		localPath = `%LOCALAPPDATA%\passh\`
		if dirErr := os.MkdirAll(`%LOCALAPPDATA%\passh\`, os.ModePerm); dirErr != nil {
			return "", fmt.Errorf("could not create directory: %w", dirErr)
		}
	} else {
		localPath = `~/.local/share/passh/`
		if dirErr := os.MkdirAll(`~/.local/share/passh`, os.ModePerm); dirErr != nil {
			return "", fmt.Errorf("could not create directory: %w", dirErr)
		}
	}
	return localPath, nil
}

func makeDBName() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("cannot set home directory: %w", err)
	}

	hash := sha1.New()
	hash.Write([]byte(homeDir))
	dbName := hex.EncodeToString(hash.Sum(nil)) + ".db"

	return dbName, nil
}
