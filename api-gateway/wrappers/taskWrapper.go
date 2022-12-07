package wrappers

import (
	"api-gateway/service"
	"context"
	"strconv"

	"github.com/afex/hystrix-go/hystrix"
	"go-micro.dev/v4/client"
)

func NewTask(id uint64, name string) *service.TaskModel {
	return &service.TaskModel{
		Id:         id,
		Title:      name,
		Content:    "out of time",
		StartTime:  1000,
		Status:     0,
		CreateTime: 1000,
		UpdateTime: 1000,
	}
}

func DefaultTasks(res interface{}) {
	models := make([]*service.TaskModel, 0)

	for i := 0; i < 10; i++ {
		models = append(models, NewTask(uint64(i), "降级"+strconv.Itoa(20+int(i))))
	}
	result := res.(*service.TaskListResponse)
	result.TaskList = models

}

type TaskWrapper struct {
	client.Client
}

func (wrapper *TaskWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	cmdName := req.Service() + "." + req.Endpoint()
	config := hystrix.CommandConfig{
		Timeout:                30000,
		RequestVolumeThreshold: 20,
		ErrorPercentThreshold:  50,
		SleepWindow:            5000,
	}

	hystrix.ConfigureCommand(cmdName, config)
	return hystrix.Do(cmdName, func() error {
		return wrapper.Client.Call(ctx, req, rsp)
	}, func(err error) error {
		return err
	})
}

func NewTaskWrapper(c client.Client) client.Client {
	return &TaskWrapper{c}
}
