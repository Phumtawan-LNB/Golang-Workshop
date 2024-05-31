package controllers

import (
	"clean/modules/entities"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type userController struct {
	userServ entities.UserService
}

func NewWeatherController(r fiber.Router, userServ entities.UserService) {
	controllers := &userController{
		userServ: userServ,
	}
	r.Post("/create", controllers.UserCreated)
	r.Get("/readed", controllers.UserReaded)
	r.Put("/update", controllers.UserUpdate)
	r.Delete("/delete", controllers.UserDelete)
}

func (obj *userController) UserCreated(c *fiber.Ctx) error {
	command := entities.UserCreatedCommand{}
	var uid = uuid.NewString()
	err := c.BodyParser(&command)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.ErrBadRequest.Message,
			"status_code": fiber.StatusBadRequest,
			"message":     "error, request body do not math",
			"result":      nil})
	}
	command.ID = uid
	id, err := obj.userServ.UserCreated(&command)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":      fiber.ErrUnprocessableEntity.Message,
			"status_code": fiber.StatusUnprocessableEntity,
			"message":     "error, cant process this entity",
			"result":      nil})
	}
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message":     "user create success",
		"id":          id,
		"status":      "OK",
		"status_code": fiber.StatusOK,
	})
}

func (obj *userController) UserReaded(c *fiber.Ctx) error {
	command := entities.UserReadedCommand{}

	err := c.BodyParser(&command)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.ErrBadRequest.Message,
			"status_code": fiber.StatusBadRequest,
			"message":     "error, request body do not math",
			"result":      nil})
	}
	data, user, err := obj.userServ.UserReaded(&command)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":      fiber.ErrUnprocessableEntity.Message,
			"status_code": fiber.StatusUnprocessableEntity,
			"message":     "error, cant process this entity",
			"result":      nil})
	}

	result := make(map[string]map[string]interface{})
	for _, weather := range data {
		weatherData := entities.WeatherData{
			Weather_id: weather.Weather_id,
			Details: map[string]interface{}{
				"name":     weather.Weather_name,
				"quantity": weather.Quantity,
			},
		}
		result[weather.Weather_id] = weatherData.Details
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "weather search success",
		"data": fiber.Map{
			"name":       user.Name,
			"lat":        user.Lat,
			"long":       user.Long,
			"age":        user.Age,
			"first_name": user.First_name,
			"last_name":  user.Last_name,
		},
		"history":     result,
		"status":      "OK",
		"status_code": fiber.StatusOK,
	})
}

func (obj *userController) UserUpdate(c *fiber.Ctx) error {
	command := entities.UserUpdateCommand{}

	err := c.BodyParser(&command)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.ErrBadRequest.Message,
			"status_code": fiber.StatusBadRequest,
			"message":     "error, request body do not math",
			"result":      nil})
	}
	id, err := obj.userServ.UserUpdate(&command)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":      fiber.ErrUnprocessableEntity.Message,
			"status_code": fiber.StatusUnprocessableEntity,
			"message":     "error, cant process this entity",
			"result":      nil})
	}
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "user update success",
		"id":      id,
		"data": fiber.Map{
			"name":       command.Name,
			"lat":        command.Lat,
			"long":       command.Long,
			"age":        command.Age,
			"first_name": command.First_name,
			"last_name":  command.Last_name,
		},
		"status":      "OK",
		"status_code": fiber.StatusOK,
	})
}

func (obj *userController) UserDelete(c *fiber.Ctx) error {
	command := entities.UserDeleteCommand{}
	err := c.BodyParser(&command)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.ErrBadRequest.Message,
			"status_code": fiber.StatusBadRequest,
			"message":     "error, request body do not math",
			"result":      nil})
	}
	id, err := obj.userServ.UserDelete(&command)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":      fiber.ErrUnprocessableEntity.Message,
			"status_code": fiber.StatusUnprocessableEntity,
			"message":     "error, cant process this entity",
			"result":      nil})
	}
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message":     "user delete success",
		"id":          id,
		"status":      "OK",
		"status_code": fiber.StatusOK,
	})
}
