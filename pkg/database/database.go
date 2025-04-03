package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

func ConnectToDB() (*gorm.DB, error) {
	_ = godotenv.Load(".env")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s sslmode=disable",
		"localhost",
		"5432",
		os.Getenv("PASSH_USER"),
		os.Getenv("PASSH_PASS"),
		"postgres")

	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		return nil, dbErr
	}

	return db, nil
}

func WizardPasswordSet(input string) (*gorm.DB, error) {
	_ = godotenv.Load(".env")

	dsn := fmt.Sprintf("host=%s port=%s database=%s sslmode=disable",
		"localhost",
		"5432",
		"postgres")

	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		return nil, dbErr
	}

	pwQuery := fmt.Sprintf(`ALTER USER %q WITH PASSWORD '%s';`, "postgres", input)
	if pwErr := db.Exec(pwQuery).Error; pwErr != nil {
		return nil, pwErr
	}

	return db, nil
}
