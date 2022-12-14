package handlers

import (
	"api-gateway/pkg/utils"
	"context"
	"proto/microtask"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetTaskList(ginCtx *gin.Context) {
	var taskReq microtask.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	taskService := ginCtx.Keys["taskService"].(microtask.TaskService)

	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)

	//Call function
	taskRes, err := taskService.GetTaskList(context.Background(), &taskReq)
	if err != nil {
		PanicIfTaskError(err)
	}
	ginCtx.JSON(200, gin.H{
		"data": gin.H{
			"task":  taskRes.TaskList,
			"count": taskRes.Count,
		},
	})
}

func CreateTaskList(ginCtx *gin.Context) {
	var taskReq microtask.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	taskService := ginCtx.Keys["taskService"].(microtask.TaskService)

	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)

	//Call function
	taskRes, err := taskService.CreateTask(context.Background(), &taskReq)
	if err != nil {
		PanicIfTaskError(err)
	}
	ginCtx.JSON(200, gin.H{
		"data": taskRes.TaskDetail,
		"msg":  "Succesfully Created",
	})
}

func GetTaskDetail(ginCtx *gin.Context) {
	var taskReq microtask.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	taskService := ginCtx.Keys["taskService"].(microtask.TaskService)

	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	//ParseToken return Uid
	taskReq.Uid = uint64(claim.Id)

	id, _ := strconv.Atoi(ginCtx.Param("id"))
	taskReq.Id = uint64(id)

	//Call function
	taskRes, err := taskService.GetTask(context.Background(), &taskReq)
	PanicIfTaskError(err)

	ginCtx.JSON(200, gin.H{
		"data": taskRes.TaskDetail,
	})
}

func UpdateTask(ginCtx *gin.Context) {
	var taskReq microtask.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	taskService := ginCtx.Keys["taskService"].(microtask.TaskService)

	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)

	id, _ := strconv.Atoi(ginCtx.Param("id"))
	taskReq.Id = uint64(id)

	//Call function
	taskRes, err := taskService.UpdateTask(context.Background(), &taskReq)

	PanicIfTaskError(err)

	ginCtx.JSON(200, gin.H{
		"data": taskRes.TaskDetail,
	})
}
func DeleteTask(ginCtx *gin.Context) {
	var taskReq microtask.TaskRequest
	PanicIfTaskError(ginCtx.Bind(&taskReq))
	taskService := ginCtx.Keys["taskService"].(microtask.TaskService)

	claim, _ := utils.ParseToken(ginCtx.GetHeader("Authorization"))
	taskReq.Uid = uint64(claim.Id)
	id, _ := strconv.Atoi(ginCtx.Param("id"))
	taskReq.Id = uint64(id)

	//Call function
	taskRes, err := taskService.DeleteTask(context.Background(), &taskReq)
	PanicIfTaskError(err)

	ginCtx.JSON(200, gin.H{
		"data": taskRes.TaskDetail,
	})
}
