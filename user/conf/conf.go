/*
 * @Author: kxxx googglexhx@gmail.com
 * @Date: 2022-11-18 10:21:49
 * @LastEditors: kxxx googglexhx@gmail.com
 * @LastEditTime: 2022-12-04 00:43:00
 * @FilePath: /micro-todolist/user/conf/conf.go
 * @Description: Configuration for RabbitMQ and MySQL
 *
 * Copyright (c) 2022 by kxxx googglexhx@gmail.com, All Rights Reserved.
 */

package conf

//组件配置
import (
	"fmt"
	"strings"
	"user/model"

	ini "gopkg.in/ini.v1"
)

var (
	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	Queue         string
	QueueProtocol string
	QueueUser     string
	QueuePassWord string
	QueueHost     string
	QueuePort     string
)

func Init() {
	//加载配置文件
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		fmt.Println("读取配置文件错误，请检查配置文件路径\nErrCode:", err)
	}

	//连接数据库
	LoadDatabaseConfig(file)
	pathMySQL := strings.Join([]string{DbUser, ":", DbPassWord, "@tcp(", DbHost, ":", DbPort, ")/", DbName, "?charset=utf8mb4&parseTime=true&loc=Local"}, "")
	model.Database(pathMySQL)

	//连接RabbitMQ
	// LoadMQConfig(file)
	// pathMQ := strings.Join([]string{QueueProtocol, "://", QueueUser, ":", QueuePassWord, "@", QueueHost, ":", QueuePort, "/"}, "")
	// model.MessageQueue(pathMQ)

}

func LoadDatabaseConfig(file *ini.File) {

	Db = file.Section("databse").Key("Db").String()
	DbHost = file.Section("databse").Key("DbHost").String()
	DbPort = file.Section("databse").Key("DbPort").String()
	DbUser = file.Section("databse").Key("DbUser").String()
	DbPassWord = file.Section("databse").Key("DbPassWord").String()
	DbName = file.Section("databse").Key("DbName").String()

}

// func LoadMQConfig(file *ini.File) {

// 	Queue = file.Section("messageQueue").Key("Queue").String()
// 	QueueProtocol = file.Section("messageQueue").Key("QueueProtocol").String()
// 	QueueUser = file.Section("messageQueue").Key("QueueUser").String()
// 	QueuePassWord = file.Section("messageQueue").Key("QueuePassWord").String()
// 	QueueHost = file.Section("messageQueue").Key("QueueHost").String()
// 	QueuePort = file.Section("messageQueue").Key("QueuePort").String()
// }
