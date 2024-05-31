package servers

import (
	_controllers "clean/modules/weathers/controllers"
	_weatherRepository "clean/modules/weathers/repositories"
	_weatherHandlerConsumer "clean/modules/consumer/handler"
	_weatherUsecasesConsumer "clean/modules/consumer/usecases"
	_weatherHandlerProducer "clean/modules/producer/handler"
	_weatherUsecasesProducer "clean/modules/producer/usecases"
	_weatherUsecases "clean/modules/weathers/usecases"
	_redis "clean/pkg/databases/redis"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) MapHandlers() error {
	// Group a version
	v1 := s.App.Group("/v1")

	weatherGroup := v1.Group("/weathers")
	weatherRepository := _weatherRepository.NewWeatherRepository(s.Db, _redis.InitRedis(s.Cfg))
	eventProducer := _weatherHandlerProducer.NewEventProducer(s.Producer)
	weatherProducer := _weatherUsecasesProducer.NewWeatherProducer(eventProducer)
	WeatherUsecase := _weatherUsecases.NewWeatherService(weatherProducer, weatherRepository)
	_controllers.NewWeatherController(weatherGroup, WeatherUsecase)

	weatherEventHandler := _weatherUsecasesConsumer.NewWeatherHandler(weatherRepository)
	weatherConsumerHandler := _weatherHandlerConsumer.NewConsumerHandler(weatherEventHandler)
	s.ConsumerHandler = weatherConsumerHandler

	// End point not found response
	s.App.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.ErrInternalServerError.Code).JSON(fiber.Map{
			"status":      fiber.ErrInternalServerError.Message,
			"status_code": fiber.ErrInternalServerError.Code,
			"message":     "error, end point not found",
			"result":      nil,
		})
	})

	return nil
}
