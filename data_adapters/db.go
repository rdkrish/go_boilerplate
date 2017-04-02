package data_adapters

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/hashicorp/errwrap"
	"go_boilerplate/config"
)

var db *gorm.DB

// Init for database
func Init() {
	c := config.GetConfig()
	connectionString := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s",
		c.GetString("DB_HOST"), c.GetString("DB_USER"), c.GetString("DB_NAME"),
		c.GetString("DB_SSL_MODE"), c.GetString("DB_PASSWORD"))
	if c.GetString("CONFIG") == "test" {
		connectionString = c.GetString("DB_PATH")
	}
	var err error
	db, err = gorm.Open(c.GetString("DB_TYPE"), connectionString)
	if errwrap.Contains(err, "credentials") {
		panic("Can't connect to database, check config!")
		// TODO handle other errors
	}
}

// GetDB to expose db object
func GetDB() *gorm.DB {
	return db
}