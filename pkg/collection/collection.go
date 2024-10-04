package collection

import (
	"log"

	"github.com/mclacore/passh/pkg/database"
	"gorm.io/gorm"
)

type Collection struct {
	gorm.Model
	ColName string `gorm:"unique"`
}

var collection Collection

func automigrateDB() {
	db, err := database.ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&Collection{})
	if err != nil {
		log.Fatal(err)
	}
}

func CreateCollection(db *gorm.DB, col Collection) error {
	automigrateDB()
	result := db.Create(&col)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetCollection(db *gorm.DB, colName string) (*Collection, error) {
	result := db.Where(&Collection{ColName: colName}).Find(&collection)
	if result.Error != nil {
		return nil, result.Error
	}
	return &collection, nil
}

func ListCollections(db *gorm.DB) (*[]Collection, error) {
	var collections []Collection

	result := db.Select("col_name").
		Order("col_name asc").
		Find(&collection)
	if result.Error != nil {
		return nil, result.Error
	}
	return &collections, nil
}

func DeleteCollection(db *gorm.DB, colName string) error {
	// need to loop through all items in collection first, delete them all, then delete collection
	result := db.Where("col_name = ?", colName).Delete(&collection)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
