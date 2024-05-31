package repositories

import (
	"clean/modules/entities"
	"clean/modules/logs"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type weatherRepository struct {
	db          *gorm.DB
	redisClient *redis.Client
}

func NewWeatherRepository(db *gorm.DB, redisClient *redis.Client) entities.WeatherRepository {
	db.Table("weather").AutoMigrate(&entities.Weather{})
	db.Table("weather_user_auth").AutoMigrate(&entities.WeatherUserAuth{})
	return weatherRepository{db, redisClient}
}

func (obj weatherRepository) Create(user entities.WeatherUserAuth) error {
	logs.Info(fmt.Sprintf("Function is Called: Create"))
	return obj.db.Table("weather_user_auth").Save(user).Error
}

func (obj weatherRepository) Save(weather entities.Weather) error {
	logs.Info(fmt.Sprintf("Function is Called: Save"))
	return obj.db.Table("weather").Save(weather).Error
}

func (obj weatherRepository) CheckId(weatherSearch entities.WeatherSearch) (user entities.WeatherUserAuth, err error) {
	logs.Info(fmt.Sprintf("Function is Called: CheckId"))
	logs.Debug(fmt.Sprintf("Data: %v", weatherSearch.ID))
	err = obj.db.Table("weather_user_auth").Where("id = ?", weatherSearch.ID).Limit(30).First(&user).Error
	if err != nil {
		logs.Error(fmt.Sprintf("Function Search Error, weatherSearch.ID: %v", weatherSearch.ID))
		return user, err
	}
	return user, nil
}

func (obj weatherRepository) Search(weatherSearch entities.WeatherSearch) (weather entities.Weather, err error) {
	logs.Info(fmt.Sprintf("Function is Called: Search"))
	logs.Debug(fmt.Sprintf("Data: %v", weatherSearch.Name))

	key := fmt.Sprintf("%v%v", weatherSearch.ID, weatherSearch.Name)
	//Redis Get
	weatherJson, err := obj.redisClient.Get(context.Background(), key).Result()
	if err == nil && weatherJson != "" {
		err = json.Unmarshal([]byte(weatherJson), &weather)
		if err != nil {
			logs.Error(err)
			return weather, nil
		}
		logs.Debug("--- On Redis ---")
	} else {
		err = obj.db.Table("weather").Where("name = ?", weatherSearch.Name).Limit(30).First(&weather).Error
		if err != nil {
			logs.Error(fmt.Sprintf("Function Search Error, weatherSearch.Name: %v", weatherSearch.Name))
			return weather, err
		}
		//Redis Set
		data, err := json.Marshal(weather)
		if err != nil {
			logs.Error(err)
			return weather, err
		}
		err = obj.redisClient.Set(context.Background(), key, string(data), time.Second*10).Err()
		if err != nil {
			logs.Error(err)
			return weather, err
		}
		logs.Debug("--- On Database ---")
		logs.Debug(fmt.Sprintf("%v", weather))
	}

	return weather, nil
}

func (obj weatherRepository) FindAll() (weather []entities.Weather, err error) {
	logs.Info(fmt.Sprintf("Function is Called: FindAll"))
	err = obj.db.Table("weather").Find(&weather).Error
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return weather, nil
}

func (obj weatherRepository) Update(id string, WeatherUser entities.Weather) error {
	logs.Info(fmt.Sprintf("Function is Called: Update"))
	logs.Debug(fmt.Sprintf("Data: %v", id))
	return obj.db.Table("weather").Where("id = ?", id).Updates(WeatherUser).Error
}

func (obj weatherRepository) UserUpdate(id string, user entities.WeatherUserAuth) error {
	logs.Info(fmt.Sprintf("Function is Called: UserUpdate"))
	logs.Debug(fmt.Sprintf("Data: %v", id))
	return obj.db.Table("weather_user_auth").Where("id = ?", id).Updates(user).Error
}

func (obj weatherRepository) Delete(id string) error {
	logs.Info(fmt.Sprintf("Function is Called: Delete"))
	logs.Debug(fmt.Sprintf("Data: %v", id))
	return obj.db.Table("weather_user_auth").Where("id = ?", id).Delete(&entities.Weather{}).Error
}

func (obj weatherRepository) WeatherDelete(id string) error {
	logs.Info(fmt.Sprintf("Function is Called: WeatherDelete"))
	logs.Debug(fmt.Sprintf("Data: %v", id))
	return obj.db.Table("weather").Where("id = ?", id).Delete(&entities.Weather{}).Error
}
