package login

import (
	"log"

	"github.com/mclacore/passh/pkg/collection"
	"github.com/mclacore/passh/pkg/database"
	"gorm.io/gorm"
)

type LoginItem struct {
	gorm.Model
	ItemName     string
	Username     string
	Password     string
	URL          string
	CollectionID int
	Collection   collection.Collection
}

var loginItem LoginItem

func automigrateDB() {
	db, err := database.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&LoginItem{})
	if err != nil {
		log.Fatal(err)
	}
}

func CreateLoginItem(db *gorm.DB, item LoginItem) error {
	automigrateDB()
	result := db.Create(&item)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetLoginItem(db *gorm.DB, itemName string, colId int) (*LoginItem, error) {
	result := db.Where(&LoginItem{ItemName: itemName, CollectionID: colId}).Find(&loginItem)
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

func ListLoginItems(db *gorm.DB, colId int) (*[]LoginItem, error) {
	var loginItems []LoginItem

	// add listing by item-name
	result := db.Select("item_name").
		Where(&LoginItem{CollectionID: colId}).
		Order("item_name asc").
		Find(&loginItems)
	if result.Error != nil {
		return nil, result.Error
	}

	return &loginItems, nil
}

func DeleteLoginItem(db *gorm.DB, itemName string, colId int) error {
	result := db.Where(&LoginItem{ItemName: itemName, CollectionID: colId}).Delete(&loginItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func AssignCollection(db *gorm.DB, itemName, colName string) error {
	var col collection.Collection

	colId := col.ID

	result := db.Where(&LoginItem{ItemName: itemName}).Set("collection_id = ?", colId)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
