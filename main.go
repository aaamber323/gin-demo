package main

import (
	"gin-demo/dao"
	"gin-demo/model"
	router "gin-demo/routers"
)

func main() {
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	defer dao.DB.Close()
	// 模型绑定
	dao.DB.AutoMigrate(&model.Todo{})
	r := router.SetupRouter()
	r.Run()
}
