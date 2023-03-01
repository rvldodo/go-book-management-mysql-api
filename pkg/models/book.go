package models

import (
	"github.com/dodo/go-book-management-mysql/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title string `gorm:""json:"title"`
	Author string `json:"author"`
	Publisher string `json:"publisher"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}