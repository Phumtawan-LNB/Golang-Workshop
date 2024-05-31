package usecases

import (
	"clean/modules/entities"
	"clean/modules/logs"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/google/uuid"
)

type weatherService struct {
	eventProducer entities.WeatherProducer
	weatherRepo   entities.WeatherRepository
}

func (obj weatherService) WeatherDelete(command entities.WeatherDeleteEvent) error {
	logs.Info(fmt.Sprintf("Function is Called: WeatherDelete"))
	logs.Debug(fmt.Sprintf("Data: Command.ID: %s", command.ID))
	err := obj.weatherRepo.WeatherDelete(command.ID)
	if err != nil {
		logs.Error(err)
		return logs.NewUnexpectedError()
	}
	return nil
}

func (obj weatherService) WeatherUpdate(command entities.WeatherUpdateEvent) (weather []entities.Weather, err error) {
	logs.Info(fmt.Sprintf("Function is Called: WeatherUpdate"))
	logs.Debug(fmt.Sprintf("Data: Command.Key: %s", command.Key))
	if command.Key != "9O1hb4rz7oPIvorzTd3N" {
		logs.Error("Key Not Found")
		return weather, logs.NewNotFoundError("Key Not Found")
	}
	data, err := obj.weatherRepo.FindAll()
	if err != nil {
		logs.Error(err)
		return weather, logs.NewNotFoundError("weather record not found")
	}
	var count = 0
	var weatherResponse entities.WeatherResponse
	// Updates the weather in the database to the current weather at the time the function is called.
	// Loop by weather_id to update current weather
	for _, item := range data {
		url := fmt.Sprintf("https://weatherapi-com.p.rapidapi.com/current.json?q=%.1f,%.1f", item.Lat, item.Long)

		req, _ := http.NewRequest("GET", url, nil)

		req.Header.Add("X-RapidAPI-Key", "37a42ccbdamsh555ce2e9a0246d9p1a6b62jsnea84bcfa516a")
		req.Header.Add("X-RapidAPI-Host", "weatherapi-com.p.rapidapi.com")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := io.ReadAll(res.Body)

		if err = json.Unmarshal(body, &weatherResponse); err != nil {
			logs.Error(err)
			return
		}
		logs.Debug(fmt.Sprintf("%+v", weatherResponse))
		updatedWeather := entities.Weather{
			Temp:     weatherResponse.Current.Temp_c,
			Wind_dir: weatherResponse.Current.Wind_dir,
			Wind_kph: weatherResponse.Current.Wind_kph,
			UV:       weatherResponse.Current.UV,
		}
		if err := obj.weatherRepo.Update(item.ID, updatedWeather); err != nil {
			logs.Error(err)
			return weather, err
		}
		count++
	}
	logs.Debug(fmt.Sprintf("count of weatherUpdate: %v", count))
	lastData, err := obj.weatherRepo.FindAll()
	if err != nil {
		logs.Error(err)
		return weather, logs.NewNotFoundError("weather record not found")
	}
	return lastData, nil
}

func (obj weatherService) WeatherSearch(command entities.WeatherSearchsEvent) (weather entities.Weather, err error) {
	logs.Info(fmt.Sprintf("Function is Called: WeatherSearch"))
	logs.Debug(fmt.Sprintf("Data: Command.ID, Command.Name: %s, %s", command.ID, command.Name))
	weatherSearch := entities.WeatherSearch{
		ID:   command.ID,
		Name: command.Name,
	}
	userData, err := obj.weatherRepo.CheckId(weatherSearch)
	logs.Debug(fmt.Sprintf("%+v", userData))
	if err != nil {
		logs.Error(err)
		return weather, logs.NewNotFoundError("user id record not found in weather_user_auth")
	}
	data, err := obj.weatherRepo.Search(weatherSearch)
	if err != nil {
		logs.Error(err)
		return weather, logs.NewNotFoundError("weather record not found")
	}
	event := &entities.WeatherSearchEvent{
		User_id:      userData.ID,
		User_name:    userData.Name,
		Weather_id:   data.ID,
		Weather_name: data.Name,
	}
	if err = obj.eventProducer.SearchEvent(event); err != nil {
		logs.Error(err)
	}
	return data, nil
}

func (obj weatherService) WeatherCreate(command entities.WeatherCreateEvent) (weather entities.Weather, err error) {
	logs.Info(fmt.Sprintf("Function is Called: WeatherCreate"))
	logs.Debug(fmt.Sprintf("Data: Command.Lat, Command.Long: %v, %v", command.Lat, command.Long))
	var uid = uuid.NewString()

	url := fmt.Sprintf("https://weatherapi-com.p.rapidapi.com/current.json?q=%.1f,%.1f", command.Lat, command.Long)

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("X-RapidAPI-Key", "37a42ccbdamsh555ce2e9a0246d9p1a6b62jsnea84bcfa516a")
	req.Header.Add("X-RapidAPI-Host", "weatherapi-com.p.rapidapi.com")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	var weatherResponse entities.WeatherResponse
	if err = json.Unmarshal(body, &weatherResponse); err != nil {
		logs.Error(err)
		return weather, logs.NewUnexpectedError()
	}
	weather = entities.Weather{
		ID:       uid,
		Name:     weatherResponse.Location.Region,
		Lat:      command.Lat,
		Long:     command.Long,
		Country:  weatherResponse.Location.Country,
		Temp:     weatherResponse.Current.Temp_c,
		Wind_dir: weatherResponse.Current.Wind_dir,
		Wind_kph: weatherResponse.Current.Wind_kph,
		UV:       weatherResponse.Current.UV,
	}
	err = obj.weatherRepo.Save(weather)
	if err != nil {
		logs.Error(err)
		return weather, logs.NewUnexpectedError()
	}
	return weather, nil
}

func NewWeatherService(eventProducer entities.WeatherProducer, weatherRepo entities.WeatherRepository) entities.WeatherService {
	return weatherService{eventProducer: eventProducer, weatherRepo: weatherRepo}
}
