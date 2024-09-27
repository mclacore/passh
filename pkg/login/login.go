package login

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type LoginItem struct {
	gorm.Model
	LoginItem string
	Username  string
	Password  string
	URL       string
}

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

	// itemID := newLoginItem.ItemName
	// item, err := getLoginItem(db, itemID)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("Found login item:", item)
	//
	// item.Username = "newUsername"
	// err = updateLoginItem(db, item)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Println("Updated login item:", item)
	//
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
	var login LoginItem
	result := db.Where(&LoginItem{LoginItem: itemName}).Find(&login)
	if result.Error != nil {
		return nil, result.Error
	}
	return &login, nil
}

func UpdateLoginItem(db *gorm.DB, loginItem *LoginItem) error {
	result := db.Save(loginItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeleteLoginItem(db *gorm.DB, loginItem *LoginItem) error {
	result := db.Delete(loginItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
