package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
func ConectToDb() error{
	var err error

	dsn :=  os.Getenv("MYSQL_DB_URL")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("faild to coneact to database ...")
		return err
	} else {
		fmt.Println("connect db ...")
	}	
	return err
}
