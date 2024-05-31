package entities

import "reflect"

type Event interface {
}

type EventProducer interface {
	Produce(event Event) error
}

type EventHandler interface {
	Handle(topic string, evenBytes []byte)
}

type WeatherProducer interface{
	SearchEvent(search *WeatherSearchEvent) error
}

var Topics = []string{
	reflect.TypeOf(UserUpdateEvent{}).Name(),
	reflect.TypeOf(UserDeleteEvent{}).Name(),
	reflect.TypeOf(UserAuthEvent{}).Name(),
	reflect.TypeOf(WeatherSearchEvent{}).Name(),
}

type UserAuthEvent struct {
	ID   string
	Name string
}

type UserUpdateEvent struct {
	ID   string
	Name string
}

type UserDeleteEvent struct {
	ID string
}

type WeatherSearchEvent struct {
	User_id      string
	User_name    string
	Weather_id   string
	Weather_name string
}
