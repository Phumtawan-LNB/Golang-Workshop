package usecases

import (
	"clean/modules/entities"
	"clean/modules/logs"
	"fmt"
)

type weatherProducer struct {
	eventProducer entities.EventProducer
}

// SearchEvent implements entities.WeatherProducer.
func (obj *weatherProducer) SearchEvent(search *entities.WeatherSearchEvent) error {
	logs.Info(fmt.Sprintf("Events is Called: SearchEvent"))
	logs.Debug(fmt.Sprintf("Data: %+v", search))
	event := entities.WeatherSearchEvent{
		User_id: search.User_id,
		User_name: search.User_name,
		Weather_id: search.Weather_id,
		Weather_name: search.Weather_name,
	}
	return obj.eventProducer.Produce(event)

}

func NewWeatherProducer(eventProducer entities.EventProducer) entities.WeatherProducer {
	return &weatherProducer{eventProducer: eventProducer}
}
