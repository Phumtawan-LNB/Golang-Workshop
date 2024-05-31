package usecases

import (
	"clean/modules/entities"
	"clean/modules/logs"
	"encoding/json"
	"fmt"
	"reflect"
)

type weatherEventHandler struct {
	weatherRepo entities.WeatherRepository
}

func NewWeatherHandler(weatherRepo entities.WeatherRepository) entities.EventHandler {
	return weatherEventHandler{weatherRepo}
}

func (obj weatherEventHandler) Handle(topic string, eventBytes []byte) {
	switch topic {
	case reflect.TypeOf(entities.UserAuthEvent{}).Name():
		logs.Info(fmt.Sprintf("Event is Called: UserAuthEvent"))
		event := &entities.UserAuthEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			logs.Error(err)
			return
		}
		user := entities.WeatherUserAuth{
			ID:   event.ID,
			Name: event.Name,
		}
		err = obj.weatherRepo.Create(user)
		if err != nil {
			logs.Error(err)
			return
		}
		logs.Debug(fmt.Sprintf("Data: Event.ID, Event.Name: %s, %s", event.ID, event.Name))
	case reflect.TypeOf(entities.UserDeleteEvent{}).Name():
		logs.Info(fmt.Sprintf("Event is Called: UserDeleteEvent"))
		event := &entities.UserDeleteEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			logs.Error(err)
			return
		}
		err = obj.weatherRepo.Delete(event.ID)
		if err != nil {
			logs.Error(err)
			return
		}
		logs.Debug(fmt.Sprintf("Data: Event.ID: %s", event.ID))
	case reflect.TypeOf(entities.UserUpdateEvent{}).Name():
		logs.Info(fmt.Sprintf("Event is Called: UserUpdateEvent"))
		event := &entities.UserUpdateEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			logs.Error(err)
			return
		}
		user := entities.WeatherUserAuth{
			Name: event.Name,
		}
		err = obj.weatherRepo.UserUpdate(event.ID, user)
		if err != nil {
			logs.Error(err)
			return
		}
		logs.Debug(fmt.Sprintf("Data: Event.ID, Event.Name: %s, %s", event.ID, event.Name))
	default:
		//log.Println("event cant get topic on event handler")
	}

}
