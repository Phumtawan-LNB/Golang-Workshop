package usecases

import (
	"clean/modules/entities"
	"clean/modules/logs"
	"fmt"
)

type userService struct {
	userRepo      entities.UserRepository
	eventProducer entities.UserProducer
}

func NewUserService(eventProducer entities.UserProducer, userRepo entities.UserRepository) entities.UserService {
	return &userService{eventProducer: eventProducer, userRepo: userRepo}
}

func (obj *userService) UserCreated(command *entities.UserCreatedCommand) (id string, err error) {
	logs.Info(fmt.Sprintf("Function is Called: UserCreated"))
	logs.Debug(fmt.Sprintf("Data: Command.Lat, Command.Long: %v, %v", command.Lat, command.Long))

	user := &entities.User{
		ID:         command.ID,
		Name:       command.Name,
		Lat:        command.Lat,
		Long:       command.Long,
		Age:        command.Age,
		First_name: command.First_name,
		Last_name:  command.Last_name,
	}
	err = obj.userRepo.Create(user)
	if err != nil {
		logs.Error(err)
		return id, logs.NewUnexpectedError()
	}

	event := &entities.UserAuthEvent{
		ID:   command.ID,
		Name: command.Name,
	}
	if err = obj.eventProducer.CreateEvent(event); err != nil {
		logs.Error(err)
	}

	logs.Debug(fmt.Sprintf("Data on Event: %v", event))
	return event.ID, nil
}

func (obj *userService) UserReaded(command *entities.UserReadedCommand) (history []entities.UserHistory, user *entities.User, err error) {
	logs.Info(fmt.Sprintf("Function is Called: UserReaded"))
	logs.Debug(fmt.Sprintf("Data: Command.User_id: %v", command.User_id))
	user, err = obj.userRepo.FindById(command.User_id)
	if err != nil {
		logs.Error(err.Error())
		return
	}
	logs.Debug(fmt.Sprintf("FindById: %v", user))

	history, err = obj.userRepo.Readed(command.User_id)
	if err != nil {
		logs.Error(err)
		return
	}
	logs.Debug(fmt.Sprintf("Readed: %v", history))
	return history, user, nil
}

func (obj *userService) UserUpdate(command *entities.UserUpdateCommand) (user *entities.User, err error) {
	logs.Info(fmt.Sprintf("Function is Called: UserUpdate"))
	logs.Debug(fmt.Sprintf("Data: %v", command))
	user = &entities.User{
		ID:         command.ID,
		Name:       command.Name,
		Lat:        command.Lat,
		Long:       command.Long,
		Age:        command.Age,
		First_name: command.First_name,
		Last_name:  command.Last_name,
	}
	_, err = obj.userRepo.Update(command.ID, user)
	if err != nil {
		logs.Error(err)
		return user, logs.NewUnexpectedError()
	}
	event := &entities.UserUpdateEvent{
		ID:   command.ID,
		Name: command.Name,
	}

	if err = obj.eventProducer.UpdateEvent(event); err != nil {
		logs.Error(err)
	}

	logs.Debug(fmt.Sprintf("Data on Event: %v", event))
	return user, nil
}

func (obj *userService) UserDelete(command *entities.UserDeleteCommand) (id string, err error) {
	logs.Info(fmt.Sprintf("Function is Called: UserCreated"))
	logs.Debug(fmt.Sprintf("Data: Command.Lat: %v", command))
	err = obj.userRepo.Delete(command.ID)
	if err != nil {
		logs.Error(err)
		return id, logs.NewUnexpectedError()
	}

	event := &entities.UserDeleteEvent{
		ID: command.ID,
	}

	if err = obj.eventProducer.DeleteEvent(event); err != nil {
		logs.Error(err)
	}
	logs.Debug(fmt.Sprintf("Data on Event: %v", event))
	return event.ID, nil
}
