package models

import (
	"fmt"
	u "go-contacts/utils"

	"github.com/jinzhu/gorm"
)

type Todo struct {
	gorm.Model
	Title     string      `json:"title"`
	UserID    uint        `json:"user_id"` //The user that this todo belongs to
	TodoItems []TodoItems `json:"todoItems,omitempty"`
}

type TodoItems struct {
	gorm.Model
	Content  string `json:"content"`
	Complete bool   `json:"complete"`
	TodoID   uint   `json:"todo_id"`
}

/*
 This struct function validate the required parameters sent through the http request body

returns message and true if the requirement is met
*/
func (todo *Todo) ValidateTodo() (map[string]interface{}, bool) {

	if todo.Title == "" {
		return u.Message(false, "Todo title should be on the payload"), false
	}

	if todo.UserID <= 0 {
		return u.Message(false, "User is not recognized"), false
	}

	//All the required parameters are present
	return u.Message(true, "success"), true
}

func (todo *Todo) CreateTodo() map[string]interface{} {

	if resp, ok := todo.ValidateTodo(); !ok {
		return resp
	}

	for _, item := range todo.TodoItems {

		GetDB().Create(item)
	}
	GetDB().Create(todo)

	resp := u.Message(true, "success")
	resp["todo"] = todo
	return resp
}

func GetTodos(user uint) []*Todo {

	todos := make([]*Todo, 0)
	err := GetDB().Table("todos").Where("user_id = ?", user).Find(&todos).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return todos
}

func GetTodo(todoID, userID uint) *Todo {

	todo := &Todo{}
	err := GetDB().Table("todos").Where("user_id = ? AND ID = ?", userID, todoID).Find(&todo).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println(todo)

	todoItems := make([]TodoItems, 0)
	err = GetDB().Table("todo_items").Where("todo_id = ?", todoID).Find(&todoItems).Error
	if err != nil {
		fmt.Println(err)
		return nil
	}

	fmt.Println(todoItems)

	todo.TodoItems = todoItems

	return todo
}

func DeleteTodo(todoID, userID uint) []*Todo {
	// selectedTodo := GetTodo(todoID, userID)
	todo := &Todo{}
	err := GetDB().Table("todos").Where("ID = ? AND user_id = ?", todoID, userID).Delete(todo).Error
	if err != nil {
		return nil
	}
	return GetTodos(userID)
}
