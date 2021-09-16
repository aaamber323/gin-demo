package controller

import (
	"gin-demo/model"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func CreateATodo(c *gin.Context) {
	//前端点击待办事项，点击提交，请求数据发送这里
	//1 从请求中把数据拿出来
	//2 把数据存入数据库
	//3 返回相应
	var todo model.Todo
	c.BindJSON(&todo)
	err := model.CreateAToDo(&todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "succcess",
			"data": todo,
		})
	}
}

func GetAllTodo(c *gin.Context) {
	todolist, err := model.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todolist)
	}
}

func UpdateATodo(c *gin.Context) {
	var appID int64
	PrameID, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	id, _ := strconv.Atoi(PrameID)
	appID = int64(id)
	todo, err := model.GetATodo(appID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	}
	c.BindJSON(&todo)
	err = model.UpdateATodo(todo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, todo)
	}
}

func DeleteATodo(c *gin.Context) {
	var appID int64
	ParamID, ok := c.Params.Get("id")
	if !ok {
		c.JSON(http.StatusOK, gin.H{"error": "无效的id"})
		return
	}
	id, _ := strconv.Atoi(ParamID)
	appID = int64(id)
	err := model.DeleteATodo(appID)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": err.Error()})
	} else {
		c.JSON(http.StatusOK, gin.H{
			ParamID: "deleted",
		})
	}

}
