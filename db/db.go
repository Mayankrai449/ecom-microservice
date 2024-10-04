package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func DB_Config() {
	var err error
	dsn := "host=localhost user=postgres password=admin dbname=inventory port=5432 sslmode=disable TimeZone=Asia/Kolkata"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database!")
}

func GetDB() *gorm.DB {
	return db
}

func Migrate(models ...interface{}) {
	err := db.AutoMigrate(models...)
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}
	log.Println("Database migrated successfully!")
}
