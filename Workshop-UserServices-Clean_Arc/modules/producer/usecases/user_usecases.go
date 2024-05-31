package usecases

import (
	"clean/modules/entities"
	"clean/modules/logs"
	"fmt"
)

type userProducer struct {
	eventProducer entities.EventProducer
}

// DeleteEvent implements entities.UserProducer.
func (obj *userProducer) DeleteEvent(users *entities.UserDeleteEvent) error {
	logs.Info(fmt.Sprintf("Events is Called: DeleteEvent"))
	logs.Debug(fmt.Sprintf("Data: Command.Lat, Command.Long: %v", users.ID))
	event := entities.UserDeleteEvent{
		ID:   users.ID,
	}
	return obj.eventProducer.Produce(event)
}

// UpdateEvent implements entities.UserProducer.
func (obj *userProducer) UpdateEvent(users *entities.UserUpdateEvent) error {
	logs.Info(fmt.Sprintf("Events is Called: CreateEvent"))
	logs.Debug(fmt.Sprintf("Data: Command.Lat, Command.Long: %v, %v", users.ID, users.Name))
	event := entities.UserUpdateEvent{
		ID:   users.ID,
		Name: users.Name,
	}
	return obj.eventProducer.Produce(event)
}

// CreateEvent implements entities.UserProducer.
func (obj *userProducer) CreateEvent(users *entities.UserAuthEvent) error {
	logs.Info(fmt.Sprintf("Events is Called: CreateEvent"))
	logs.Debug(fmt.Sprintf("Data: Command.Lat, Command.Long: %v, %v", users.ID, users.Name))
	event := entities.UserAuthEvent{
		ID:   users.ID,
		Name: users.Name,
	}
	return obj.eventProducer.Produce(event)

}

func NewUserProducer(eventProducer entities.EventProducer) entities.UserProducer {
	return &userProducer{eventProducer: eventProducer}
}
