package core

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"proto/microtask"
	"task/model"

	"github.com/streadway/amqp"
)

//memorandum created
//Memorandum transformed into RabbitMQ Queue
func (*TaskService) CreateTask(ctx context.Context, req *microtask.TaskRequest, resp *microtask.TaskDetailResponse) error {
	ch, err := model.MQ.Channel()
	if err != nil {
		err := errors.New("RabbitMQ Error: " + err.Error())
		log.Fatal(err)
	}
	q, _ := ch.QueueDeclare(
		"task_queue",
		true,
		false,
		false,
		false,
		nil)
	body, _ := json.Marshal(req)
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         body,
		})
	if err != nil {
		err = errors.New("RabbitMQ Publish Error: " + err.Error())
	}

	return err

}

//GetList
func (*TaskService) GetTaskList(ctx context.Context, req *microtask.TaskRequest, res *microtask.TaskListResponse) error {
	if req.Limit == 0 {
		req.Limit = 10
	}
	var taskData []model.Task
	var count int64

	err := model.DB.Offset(int(req.Start)).Limit(int(req.Limit)).Where("uid=?", req.Uid).Find(&taskData).Error
	if err != nil {
		return errors.New("mysql find error :" + err.Error())
	}

	model.DB.Model(&model.Task{}).Where("uid=?", req.Uid).Count(&count)

	var taskRes []*microtask.TaskModel
	for _, item := range taskData {
		taskRes = append(taskRes, BuildTask(item))
	}

	res.TaskList = taskRes

	res.Count = uint32(count)
	return nil
}

//GetDetail
func (*TaskService) GetTask(ctx context.Context, req *microtask.TaskRequest, res *microtask.TaskDetailResponse) error {
	taskData := model.Task{}
	model.DB.First(&taskData, req.Id)
	taskRes := BuildTask(taskData)
	res.TaskDetail = taskRes
	return nil
}

//UpdateDetail
func (*TaskService) UpdateTask(ctx context.Context, req *microtask.TaskRequest, res *microtask.TaskDetailResponse) error {
	taskData := model.Task{}
	model.DB.Model(&model.Task{}).Where("id=? AND uid=?", req.Id, req.Uid).First(&taskData)
	taskData.Title = req.Title
	taskData.Status = int(req.Status)
	taskData.Content = req.Content
	model.DB.Save(&taskData)
	res.TaskDetail = BuildTask(taskData)
	return nil
}

//DeleteDetail
func (*TaskService) DeleteTask(ctx context.Context, req *microtask.TaskRequest, res *microtask.TaskDetailResponse) error {
	model.DB.Model(&model.Task{}).Where("id=? AND uid=?", req.Id, req.Uid).Delete(&model.Task{})
	return nil
}
