package login

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type LoginItem struct {
	gorm.Model
	ItemName string `gorm:"unique"`
	Username string
	Password string
	URL      string
}

var loginItem LoginItem

func ConnectToDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func automigrateDB() {
	db, err := ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&LoginItem{})
	if err != nil {
		log.Fatal(err)
	}

	// err = deleteLoginItem(db, item)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("Deleted login item:", item)
}

func CreateLoginItem(db *gorm.DB, item LoginItem) error {
	automigrateDB()
	result := db.Create(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetLoginItem(db *gorm.DB, itemName string) (*LoginItem, error) {
	result := db.Where(&LoginItem{ItemName: itemName}).Find(&loginItem)
	if result.Error != nil {
		return nil, result.Error
	}
	return &loginItem, nil
}

func UpdateLoginItem(db *gorm.DB, loginItem *LoginItem) error {
	result := db.Save(loginItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func ListLoginItems(db *gorm.DB) (*[]LoginItem, error) {
	var loginItems []LoginItem

	result := db.Select("item_name").
		Order("item_name asc").
		Find(&loginItems)
	if result.Error != nil {
		return nil, result.Error
	}

	return &loginItems, nil
}

func DeleteLoginItem(db *gorm.DB, itemName string) error {
	result := db.Where("item_name = ?", itemName).Delete(&loginItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
