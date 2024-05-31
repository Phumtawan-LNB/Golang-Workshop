package controllers

import (
	"clean/modules/entities"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type weatherController struct {
	weatherServ entities.WeatherService
}

func NewWeatherController(r fiber.Router, weatherServ entities.WeatherService) {
	controllers := &weatherController{
		weatherServ: weatherServ,
	}
	r.Post("/create", controllers.WeatherCreate)
	r.Get("/search", controllers.WeatherSearch)
	r.Get("/update", controllers.WeatherUpdate)
	r.Delete("/delete", controllers.WeatherDelete)
}

func (obj weatherController) WeatherDelete(c *fiber.Ctx) error {
	command := entities.WeatherDeleteEvent{}

	err := c.BodyParser(&command)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.ErrBadRequest.Message,
			"status_code": fiber.StatusBadRequest,
			"message":     "error, request body do not math",
			"result":      nil})
	}
	err = obj.weatherServ.WeatherDelete(command)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":      fiber.ErrUnprocessableEntity.Message,
			"status_code": fiber.StatusUnprocessableEntity,
			"message":     "error, cant process this entity",
			"result":      nil})
	}
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message":     "weather delete success",
		"id":          command.ID,
		"status":      "OK",
		"status_code": fiber.StatusOK,
	})
}

func (obj weatherController) WeatherUpdate(c *fiber.Ctx) error {
	command := entities.WeatherUpdateEvent{}
	err := c.BodyParser(&command)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.ErrBadRequest.Message,
			"status_code": fiber.StatusBadRequest,
			"message":     "error, request body do not math",
			"result":      nil})
	}
	data, err := obj.weatherServ.WeatherUpdate(command)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":      fiber.ErrUnprocessableEntity.Message,
			"status_code": fiber.StatusUnprocessableEntity,
			"message":     "error, cant process this entity",
			"result":      nil})
	}
	result := make(map[string]map[string]interface{})

	for _, weather := range data {
		weatherMap := map[string]interface{}{
			"lat":      weather.Lat,
			"long":     weather.Long,
			"Country":  weather.Country,
			"temp":     weather.Temp,
			"Wind_dir": weather.Wind_dir,
			"wind_kph": weather.Wind_kph,
			"uv":       weather.UV,
		}
		result[weather.Name] = weatherMap
	}

	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message":     "weather search success",
		"data":        result,
		"status":      "OK",
		"status_code": fiber.StatusOK,
	})
}

func (obj weatherController) WeatherSearch(c *fiber.Ctx) error {
	command := entities.WeatherSearchsEvent{}

	err := c.BodyParser(&command)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.ErrBadRequest.Message,
			"status_code": fiber.StatusBadRequest,
			"message":     "error, request body do not math",
			"result":      nil})
	}
	data, err := obj.weatherServ.WeatherSearch(command)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":      fiber.ErrUnprocessableEntity.Message,
			"status_code": fiber.StatusUnprocessableEntity,
			"message":     "error, cant process this entity",
			"result":      nil})
	}
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "weather search success",
		"id":      data.ID,
		"data": fiber.Map{
			"name":     data.Name,
			"lat":      data.Lat,
			"long":     data.Long,
			"country":  data.Country,
			"temp":     data.Temp,
			"wind_dir": data.Wind_dir,
			"wind_kph": data.Wind_kph,
			"UV":       data.UV,
		},
		"status":      "OK",
		"status_code": fiber.StatusOK,
	})
}

func (obj weatherController) WeatherCreate(c *fiber.Ctx) error {
	command := entities.WeatherCreateEvent{}

	err := c.BodyParser(&command)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"status":      fiber.ErrBadRequest.Message,
			"status_code": fiber.StatusBadRequest,
			"message":     "error, request body do not math",
			"result":      nil})
	}
	data, err := obj.weatherServ.WeatherCreate(command)
	if err != nil {
		return c.Status(http.StatusUnprocessableEntity).JSON(fiber.Map{
			"status":      fiber.ErrUnprocessableEntity.Message,
			"status_code": fiber.StatusUnprocessableEntity,
			"message":     "error, cant process this entity",
			"result":      nil})
	}
	c.Status(fiber.StatusOK)
	return c.JSON(fiber.Map{
		"message": "weather create success",
		"id":      data.ID,
		"data": fiber.Map{
			"name":     data.Name,
			"lat":      command.Lat,
			"long":     command.Long,
			"country":  data.Country,
			"temp":     data.Temp,
			"wind_dir": data.Wind_dir,
			"wind_kph": data.Wind_kph,
			"UV":       data.UV,
		},
		"status":      "OK",
		"status_code": fiber.StatusOK,
	})
}
