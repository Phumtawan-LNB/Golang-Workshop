package entities

import "reflect"

var Topics = []string{
	reflect.TypeOf(UserUpdateEvent{}).Name(),
	reflect.TypeOf(UserDeleteEvent{}).Name(),
	reflect.TypeOf(UserAuthEvent{}).Name(),
	reflect.TypeOf(WeatherSearchEvent{}).Name(),
}

type UserProducer interface{
	CreateEvent (user *UserAuthEvent) error
	UpdateEvent (user *UserUpdateEvent) error
	DeleteEvent (user *UserDeleteEvent) error
}

type Event interface {
}

type EventHandler interface {
	Handle(topic string, evenBytes []byte)
}

type EventProducer interface {
	Produce(event Event) error
}

type WeatherSearchEvent struct {
	User_id      string
	User_name    string
	Weather_id   string
	Weather_name string
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
