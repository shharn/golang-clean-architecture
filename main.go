package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shharn/golang-clean-architecture/handler"
	"github.com/shharn/golang-clean-architecture/usecase"
	"github.com/shharn/golang-clean-architecture/repository"
)

func main() {
	todoRepository := repository.NewTodoRepository()
	todoUsecase := usecase.NewTodoUsecase(todoRepository)
	todoHandler := handler.NewTodoHandler(todoUsecase)

	r := gin.Default()
	r.GET("/todos", todoHandler.GetAll)
	r.Run(":8080")
}