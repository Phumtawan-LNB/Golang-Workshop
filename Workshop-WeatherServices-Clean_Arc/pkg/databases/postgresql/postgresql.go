package postgresql

import (
	"clean/configs"
	"clean/modules/logs"
	"clean/pkg/utils"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgreSQLDBConnection(cfg *configs.Configs) (*gorm.DB, error) {
	postgresUrl, err := utils.ConnectionUrlBuilder("postgresql", cfg)
	if err != nil {
		return nil, err
	}
	dial := postgres.Open(postgresUrl)
	db, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		logs.Error(fmt.Sprintf("error, can't connect to database, %s", err.Error()))
		return nil, err
	}

	logs.Info("postgreSQL database has been connected 🐘")
	return db, nil
}