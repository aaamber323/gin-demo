package model

import (
	"gin-demo/dao"
)

//ToDo Model
type Todo struct {
	dao.Model
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CreateAToDo(todo *Todo) (err error) {
	err = dao.DB.Create(&todo).Error
	return
}

func GetAllTodo() (todolist []*Todo, err error) {
	if err = dao.DB.Find(&todolist).Error; err != nil {
		return nil, err
	}
	return
}

func GetATodo(id int64) (todo *Todo, err error) {
	todo = new(Todo)
	if err = dao.DB.Where("id=?", id).First(&todo).Error; err != nil {
		return nil, err
	}
	return
}

func UpdateATodo(todo *Todo) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}

func DeleteATodo(id int64) (err error) {
	err = dao.DB.Where("id=?", id).Delete(&Todo{}).Error
	return
}
