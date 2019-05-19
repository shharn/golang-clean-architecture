package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/shharn/golang-clean-architecture/usecase"
)

type TodoHandler struct {
	tu usecase.TodoUsecase
}

func (h *TodoHandler) GetAll(c *gin.Context) {
	todos := h.tu.GetAllTodos()
	c.JSON(200, todos)
}

func NewTodoHandler(tu usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{tu}
}