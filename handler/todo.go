package handler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/shharn/golang-clean-architecture/model"
	"github.com/shharn/golang-clean-architecture/usecase"
)

type createTodoInput struct {
	Content string `json:"content"`
}

type updateTodoInput struct {
	Id int `json:"id"`
	Content string `json:"content"`
}

type TodoHandler struct {
	tu usecase.TodoUsecase
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var input createTodoInput
	c.Bind(&input)
	todo := model.Todo{Content: input.Content}
	err := h.tu.CreateTodo(todo)
	if err != nil {
		log.Printf("Error occured - %+v", err)
		c.Status(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	var input updateTodoInput
	c.Bind(&input)
	todo := model.Todo{Id: input.Id, Content: input.Content}
	err := h.tu.UpdateTodo(todo)
	if err != nil {
		log.Printf("Error occured - %+v", err)
		c.Status(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Printf("Error occured - %+v", err)
		c.Status(http.StatusBadRequest)
	}

	if err := h.tu.DeleteTodo(id); err != nil {
		log.Printf("Error occured -  %+v", err)
		c.Status(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}

func (h *TodoHandler) GetAllTodos(c *gin.Context) {
	todos := h.tu.GetAllTodos()
	c.JSON(200, todos)
}

func NewTodoHandler(tu usecase.TodoUsecase) *TodoHandler {
	return &TodoHandler{tu}
}