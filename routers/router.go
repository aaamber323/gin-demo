package router

import (
	controller "gin-demo/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//告诉gin框架模板的静态文件路径
	r.Static("static/", "static")

	//告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", controller.IndexHandler)

	//代办事项
	v1Group := r.Group("v1")
	{
		//添加代办事项
		v1Group.POST("/todo", controller.CreateATodo)

		//查询全部代办事项
		v1Group.GET("/todo", controller.GetAllTodo)

		//修改待办事项
		v1Group.PUT("/todo/:id", controller.UpdateATodo)

		//删除代办事项
		v1Group.DELETE("todo/:id", controller.DeleteATodo)

	}
	return r
}
