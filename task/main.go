package main

import (
	"log"
	"task/conf"
	"task/core"
	"task/service"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func main() {
	conf.Init()
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	//micro-service registed
	s := micro.NewService(
		micro.Name("todolist.service.task"),
		//Task register PORT:8083
		micro.Address("127.0.0.1:8083"),
		micro.Registry(etcdReg),
	)

	s.Init()

	if err := service.RegisterTaskServiceHandler(s.Server(), new(core.TaskService)); err != nil {
		log.Panic(err)
	}

	if err := s.Run(); err != nil {
		log.Panic(err)
	}
}
