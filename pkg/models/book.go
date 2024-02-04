package models

import (
	"bookstore/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type book struct {
	gorm.model
	Name   string `gorm:""json:"name"`
	Author string `json:"author"`
	Publications string `json:"publications"`
}

func init() {
	config.Connect()
}
