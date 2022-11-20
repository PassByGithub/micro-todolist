package model

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func Database(connstring string) {
	db, err := gorm.Open(mysql.Open(connstring), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("Cann't connect to the databse")
	}

	//开启日志功能

	sqldb, err := db.DB()

	sqldb.SetMaxIdleConns(20)
	sqldb.SetMaxOpenConns(100)
	sqldb.SetConnMaxLifetime(time.Second * 30)

	DB = db

	//?
	//sqldb.Set(`gorm:table_options`, "charset=utf8mb4").
	db.AutoMigrate(&User{})

}
