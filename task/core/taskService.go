package core

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"task/model"
	"task/service"

	"github.com/streadway/amqp"
)

//memorandum created
//Memorandum transformed into RabbitMQ Queue
func (*TaskService) CreateTask(ctx context.Context, req *service.TaskRequest, resp *service.TaskDetailResponse) error {
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
