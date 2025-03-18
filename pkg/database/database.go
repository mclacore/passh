package database

import (
	// "crypto/sha1"
	// "encoding/hex"
	"fmt"
	"os"
	// "runtime"

	// "gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
	// osPath, pathErr := setPath()
	// if pathErr != nil {
	// 	return nil, pathErr
	// }
	//
	// dbName, nameErr := makeDBName()
	// if nameErr != nil {
	// 	return nil, nameErr
	// }
	//
	// dbPath := osPath + dbName
	//
	// if _, err := os.Stat(dbPath); err != nil {
	// 	dbFile, fileErr := os.Create(dbPath)
	// 	if fileErr != nil {
	// 		return nil, fmt.Errorf("cannot create database file: %w", fileErr)
	// 	}
	// 	defer dbFile.Close()
	// }

	if os.Getenv("PASSH_DB_PORT") == "" {
		os.Setenv("PASSH_DB_PORT", "5432")
	}

	if os.Getenv("PASSH_DB_HOST") == "" {
		// prompt setup wizard
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable TimeZone=%s",
		os.Getenv("PASSH_DB_HOST"),
		os.Getenv("PASSH_DB_PORT"),
		os.Getenv("PASSH_DB_USER"),
		os.Getenv("PASSH_PASS"),
		os.Getenv("PASSH_DB_TIMEZONE"))

	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		return nil, dbErr
	}

	if os.Getenv("PASSH_DB_NAME") == "" {
		os.Setenv("PASSH_DB_NAME", "passh")
	}

	createDbCmd := fmt.Sprintf("CREATE DATABASE %s", os.Getenv("PASSH_DB_NAME"))
	db.Exec(createDbCmd)

	return db, nil
}

// func setPath() (string, error) {
// 	var localPath string
//
// 	homeDir, err := os.UserHomeDir()
// 	if err != nil {
// 		return "", fmt.Errorf("cannot set home directory: %w", err)
// 	}
//
// 	if runtime.GOOS == "windows" {
// 		localPath = fmt.Sprintf("%v/passh/data/", homeDir)
// 		if dirErr := os.MkdirAll(localPath, os.ModePerm); dirErr != nil {
// 			return "", fmt.Errorf("could not create directory: %w", dirErr)
// 		}
// 	} else {
// 		localPath = fmt.Sprintf("%v/.local/share/passh/data/", homeDir)
// 		if dirErr := os.MkdirAll(localPath, os.ModePerm); dirErr != nil {
// 			return "", fmt.Errorf("could not create directory: %w", dirErr)
// 		}
// 	}
// 	return localPath, nil
// }

// func makeDBName() (string, error) {
// 	homeDir, err := os.UserHomeDir()
// 	if err != nil {
// 		return "", fmt.Errorf("cannot get home directory: %w", err)
// 	}
//
// 	hash := sha1.New()
// 	hash.Write([]byte(homeDir))
// 	dbName := hex.EncodeToString(hash.Sum(nil)) + ".db"
//
// 	return dbName, nil
// }
