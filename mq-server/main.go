package main

import (
	"fmt"
	"mq-server/conf"
	"mq-server/service"
)

func main() {
	conf.Init()
	fmt.Println("Connecting Established!")

	forever := make(chan bool)
	service.CreateTask()
	<-forever

}
