package setup

import (
	"SchoolMgtmt/src/models"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {
	var err error
	dsn := "root:mysqlsecret@tcp(mysql:3306)/school_mgt_system?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	err2 := DB.AutoMigrate(&models.Student{}, &models.Course{})
	if err2 != nil {
		panic(err2.Error())
	}
	if err != nil {
		panic(err.Error())

	}
	fmt.Println("Success! Connected to Database")
}
func GetDB() *gorm.DB {
	return DB
}
