package database

import (
	"fmt"
	instituteEntity "gihub.com/toufiq-austcse/vNoticeboard/api/institute/entities"
	noticeEntity "gihub.com/toufiq-austcse/vNoticeboard/api/notice/entities"
	"github.com/joho/godotenv"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func SetupDBConnection() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	fmt.Println("dbName ", dbName)
	dsn := dbUser + ":@/" + dbName
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB Connection Error ", err)
		return nil
	}
	fmt.Println("Connected To DB")
	fmt.Println("Migrating....")
	db.AutoMigrate(instituteEntity.Institute{}, noticeEntity.Notice{})
	return db
}

func CloseDBConnection(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to close connection from DB")
	}
	dbSQL.Close()
}
