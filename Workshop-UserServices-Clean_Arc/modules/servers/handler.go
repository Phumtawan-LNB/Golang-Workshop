package servers

import (
	_UserHandlerProducer "clean/modules/producer/handler"
	_UserUsecaseProducer "clean/modules/producer/usecases"
	_controllers "clean/modules/users/controllers"
	_usersRepository "clean/modules/users/repositories"
	_UserUsecase "clean/modules/users/usecases"
	_userUsecaseConsumer "clean/modules/consumer/handler"
	_userUsecaseHandler "clean/modules/consumer/usecases"

	//_redis "clean/pkg/databases/redis"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) MapHandlers() error {
	// Group a version
	v1 := s.App.Group("/v1")

	// Users Group , Kafka Producer
	userGroup := v1.Group("/users")
	userRepository := _usersRepository.NewUserRepository(s.Db)
	eventProducer := _UserHandlerProducer.NewEventProducer(s.Producer)
	userProducer := _UserUsecaseProducer.NewUserProducer(eventProducer)
	userUsecase := _UserUsecase.NewUserService(userProducer, userRepository)
	_controllers.NewWeatherController(userGroup, userUsecase)

	// Kafka Consumer 
	userEventHandler := _userUsecaseHandler.NewUserHandler(userRepository)
	userConsumerHandler := _userUsecaseConsumer.NewConsumerHandler(userEventHandler)
	s.ConsumerHandler = userConsumerHandler

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
