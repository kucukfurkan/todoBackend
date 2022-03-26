package handler

import (
	"Desktop/todo-backend/model"
	"Desktop/todo-backend/service"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	CreateTodo(ctx *fiber.Ctx) error
	GetTodoElements(ctx *fiber.Ctx) error
}

type handler struct {
	service service.Service
}

var _ Handler = handler{}

func NewHandler(service service.Service) Handler {
	return handler{service: service}
}

type Response struct {
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func (h handler) GetTodoElements(c *fiber.Ctx) error {
	model, err := h.service.GetTodoElements()
	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error()})
	}

	return c.Status(200).JSON(model)
}

func (h handler) CreateTodo(c *fiber.Ctx) error {
	todo := model.SendTodoElements{}

	err := c.BodyParser(&todo)

	if err != nil {
		return c.Status(400).JSON(Response{Error: err.Error()})
	}

	err = h.service.CreateTodo(todo)

	return c.SendStatus(201)
}
