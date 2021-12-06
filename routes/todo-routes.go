package routes

import (
	"api/controllers"
)

func init() {
	register("POST", "/api/todo", controllers.AddTodo, nil)
	register("GET", "/api/todo/{id}", controllers.GetTodoById, nil)
}
