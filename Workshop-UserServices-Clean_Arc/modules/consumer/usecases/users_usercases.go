package usecases

import (
	"clean/modules/entities"
	"clean/modules/logs"
	"encoding/json"
	"fmt"
	"reflect"
)

type userEventHandler struct {
	userRepo entities.UserRepository
}

func NewUserHandler(userRepo entities.UserRepository) entities.EventHandler {
	return &userEventHandler{userRepo}
}

func (obj *userEventHandler) Handle(topic string, eventBytes []byte) {
	switch topic {
	case reflect.TypeOf(entities.WeatherSearchEvent{}).Name():
		logs.Info(fmt.Sprintf("Event is Called: WeatherSearchEvent"))
		event := &entities.WeatherSearchEvent{}
		err := json.Unmarshal(eventBytes, event)
		if err != nil {
			logs.Error(err)
			return
		}
		data, err := obj.userRepo.FindHistory(event.User_id, event.Weather_id)
		if err != nil {
			logs.Debug("check history record is empty")
		}
		if data.User_id == event.User_id && data.Weather_id == event.Weather_id {
			quantity := data.Quantity
			quantity++
			history := entities.UserHistory{
				User_id:      event.User_id,
				User_name:    event.User_name,
				Weather_id:   event.Weather_id,
				Weather_name: event.Weather_name,
				Quantity:     quantity,
			}
			err = obj.userRepo.UpdateHistory(data.User_id, data.Weather_id, &history)
			if err != nil {
				logs.Error(err)
				return
			}
		} else {
			logs.Debug(fmt.Sprintf("%v", data))
			history := entities.UserHistory{
				User_id:      event.User_id,
				User_name:    event.User_name,
				Weather_id:   event.Weather_id,
				Weather_name: event.Weather_name,
				Quantity:     1,
			}
			err = obj.userRepo.Save(&history)
			if err != nil {
				logs.Error(err)
				return
			}
		}

		logs.Debug(fmt.Sprintf("Data: %s, %s, %s, %s", event.User_id, event.Weather_id, event.User_name, event.Weather_name))
	default:
		//log.Println("event cant get topic on event handler")
	}

}
