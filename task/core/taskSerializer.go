package core

import (
	"proto/microtask"
	"task/model"
	"time"
)

func BuildTask(item model.Task) *microtask.TaskModel {
	taskModel := microtask.TaskModel{
		Id:         uint64(item.ID),
		Uid:        uint64(item.Uid),
		Title:      item.Title,
		Content:    item.Content,
		StartTime:  item.StartTime,
		Status:     int64(item.Status),
		CreateTime: item.CreatedAt.Format(time.UnixDate),
		UpdateTime: item.UpdatedAt.Format(time.UnixDate),
	}
	return &taskModel
}
