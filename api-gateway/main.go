package main

import (
	"api-gateway/weblib"
	"api-gateway/wrappers"
	taskproto "proto/microtask"
	userproto "proto/microuser"
	"time"

	"github.com/go-micro/plugins/v4/registry/etcd"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/web"
)

func main() {
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)

	userMicroService := micro.NewService(
		micro.Name("userService.client"),
		micro.WrapClient(wrappers.NewUserWrapper),
	)

	taskMicroService := micro.NewService(
		micro.Name("taskService.client"),
		micro.WrapClient(wrappers.NewTaskWrapper),
	)

	userService := userproto.NewUserService("todolist.service.user", userMicroService.Client())
	taskService := taskproto.NewTaskService("todolist.service.task", taskMicroService.Client())
	server := web.NewService(
		web.Name("httpService"),
		web.Address(":4000"),
		web.Handler(weblib.NewRouter(userService, taskService)),
		web.Registry(etcdReg),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*15),
		web.Metadata(map[string]string{"protocol": "http"}),
	)

	_ = server.Init()
	_ = server.Run()

}
