package database

import (
	"fmt"

	"github.com/mclacore/passh/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToDB() (*gorm.DB, error) {
	user, userErr := config.LoadConfigValue("auth", "username")
	if userErr != nil {
		return nil, userErr
	}

	pass, passErr := config.LoadConfigValue("auth", "persist_pass")
	if pass == "" {
		pass, passErr = config.LoadConfigValue("auth", "temp_pass")
	}

	if passErr != nil {
		return nil, passErr
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s database=%s sslmode=disable",
		"localhost",
		"5432",
		user,
		pass,
		"postgres")

	db, dbErr := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if dbErr != nil {
		return nil, dbErr
	}

	return db, nil
}

func WizardPasswordSet(input string) (*gorm.DB, error) {
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
