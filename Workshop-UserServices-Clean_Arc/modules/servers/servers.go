package servers

import (
	"clean/configs"
	"clean/modules/entities"
	"clean/modules/logs"
	"clean/pkg/utils"
	"context"
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Server struct {
	App             *fiber.App
	Cfg             *configs.Configs
	Db              *gorm.DB
	Producer        sarama.SyncProducer
	Consumer        sarama.ConsumerGroup
	ConsumerHandler sarama.ConsumerGroupHandler
}

func NewServer(cfg *configs.Configs, db *gorm.DB, producer sarama.SyncProducer, consumer sarama.ConsumerGroup) *Server {
	return &Server{
		App:      fiber.New(),
		Cfg:      cfg,
		Db:       db,
		Producer: producer,
		Consumer: consumer,
	}
}

func (s *Server) Start() {
	if err := s.MapHandlers(); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

	fiberConnURL, err := utils.ConnectionUrlBuilder("fiber", s.Cfg)
	if err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}
	go func() {
		logs.Info(fmt.Sprintf("--- User Consumer Started⚡ ---"))
		for {
			s.Consumer.Consume(context.Background(), entities.Topics, s.ConsumerHandler)
		}
	}()

	host := s.Cfg.App.Host
	port := s.Cfg.App.Port
	logs.Info(fmt.Sprintf("server has been started on %s:%s ⚡", host, port))

	if err := s.App.Listen(fiberConnURL); err != nil {
		log.Fatalln(err.Error())
		panic(err.Error())
	}

}
