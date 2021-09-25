package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Init() *gorm.DB {
	dbUser:=os.Getenv("DB_USER")
	dbName :=os.Getenv("DB_NAME")
	dsn := dbUser+":@/"+dbName
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("DB Connection Error ", err)
		return nil
	}
	DB = db
	fmt.Println("Connected To DB")
	return db
}
