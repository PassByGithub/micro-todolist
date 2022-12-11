package model

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

//全局调用DB数据库，所以要暴露出来
var DB *gorm.DB

func Database(connstring string) {
	db, err := gorm.Open(mysql.Open(connstring), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}

	sqldb, err := db.DB()

	sqldb.SetMaxIdleConns(20)
	sqldb.SetMaxOpenConns(100)
	sqldb.SetConnMaxLifetime(time.Second * 30)
	if err != nil {
		panic(err)
	}

	DB = db

	db.AutoMigrate(&Task{})
}
