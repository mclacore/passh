package collection

import (
	"log"

	"github.com/mclacore/passh/pkg/database"
	"gorm.io/gorm"
)

type Collection struct {
	gorm.Model
	Name string `gorm:"uniqueIndex"` // gorm:unique is not working? maybe add checking in create step
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
	result := db.Where(&Collection{Name: colName}).Find(&collection)
	if result.Error != nil {
		return nil, result.Error
	}
	return &collection, nil
}

func GetCollectionById(db *gorm.DB, colId int) (*Collection, error) {
	result := db.Select("name").Where("id = ?", colId).Find(&collection)		
	if result.Error != nil {
		return nil, result.Error
	}
	return &collection, nil
}

func ListCollections(db *gorm.DB) (*[]Collection, error) {
	var collections []Collection

	result := db.Select("name").
		Order("name asc").
		Find(&collections)
	if result.Error != nil {
		return nil, result.Error
	}
	return &collections, nil
}

func DeleteCollection(db *gorm.DB, colName string) error {
	result := db.Where("name = ?", colName).Delete(&collection)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
