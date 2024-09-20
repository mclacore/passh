package login

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type LoginItem struct {
	ItemName string `gorm:"primaryKey"`
	Username string
	Password string
	URL      string
}

func connectToDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func manageDB() {
	db, err := connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&LoginItem{})
	if err != nil {
		log.Fatal(err)
	}

	newLoginItem := &LoginItem{ItemName: "test", Username: "test", Password: "test", URL: "test"}
	err = createLoginItem(db, newLoginItem)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Created a new login item:", newLoginItem)

	itemName := newLoginItem.ItemName
	item, err := getLoginItem(db, itemName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Found login item:", item)

	item.Username = "newUsername"
	err = updateLoginItem(db, item)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Updated login item:", item)

	err = deleteLoginItem(db, item)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Deleted login item:", item)
}

func createLoginItem(db *gorm.DB, item LoginItem) error {
	result := db.Create(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func getLoginItem(db *gorm.DB, itemName string) (*LoginItem, error) {
	var login LoginItem
	result := db.First(&login, itemName)
	if result.Error != nil {
		return nil, result.Error
	}
	return &login, nil
}

func updateLoginItem(db *gorm.DB, loginItem *LoginItem) error {
	result := db.Save(loginItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func deleteLoginItem(db *gorm.DB, loginItem *LoginItem) error {
	result := db.Delete(loginItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
