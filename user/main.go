package main

import (
	"log"
	"user/core"
	"user/services"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
)

func main() {

	//etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	//micro-service registed
	s := micro.NewService(
		micro.Name("todolist.service.user"),
		micro.Address("127.0.0.1:8082"),
		micro.Registry(etcdReg),
	)

	s.Init()

	if err := services.RegisterUserServiceHandler(s.Server(), new(core.UserService)); err != nil {
		log.Panic(err)
	}

	if err := s.Run(); err != nil {
		log.Panic(err)
	}

}
