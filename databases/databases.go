package databases

import (
	"fibermongo/models"
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	Db *gorm.DB
)

func InitDatabase() {
	var err error

	Db, err = gorm.Open("mysql", "root:01052020@tcp(127.0.0.1:3306)/gorm?parseTime=true")

	log.Println(err)
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")
	Db.AutoMigrate(&models.User{}, &models.Post{})
	fmt.Println("Database Migrated")
}
