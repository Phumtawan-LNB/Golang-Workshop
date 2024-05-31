package repositories

import (
	"clean/modules/entities"
	"clean/modules/logs"
	"fmt"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) entities.UserRepository {
	db.Table("user").AutoMigrate(&entities.User{})
	db.Table("user_history").AutoMigrate(&entities.UserHistory{})
	return &userRepository{db}
}

func (obj *userRepository) Create(user *entities.User) error {
	logs.Info(fmt.Sprintf("Function is Called: Create"))
	return obj.db.Table("user").Save(user).Error
}

func (obj *userRepository) Save(history *entities.UserHistory) error {
	logs.Info(fmt.Sprintf("Function is Called: Save"))
	return obj.db.Table("user_history").Create(history).Error
}

func (obj *userRepository) UpdateHistory(id string, weather_id string, history *entities.UserHistory) error {
	logs.Info(fmt.Sprintf("Function is Called: UpdateHistory"))
	logs.Debug(fmt.Sprintf("Data: %v, %v", id, weather_id))
	return obj.db.Table("user_history").Where("user_id = ? and weather_id = ?", id, weather_id).Updates(history).Error
}

func (obj *userRepository) FindById(id string) (user *entities.User, err error) {
	logs.Info(fmt.Sprintf("Function is Called: FindById"))
	logs.Debug(fmt.Sprintf("Data: %v", id))
	err = obj.db.Table("user").Where("id = ?", id).First(&user).Error
	return user, err
}

func (obj *userRepository) FindHistory(id string, weather_id string) (history *entities.UserHistory, err error) {
	logs.Info(fmt.Sprintf("Function is Called: FindHistory"))
	logs.Debug(fmt.Sprintf("Data: %v, %v", id, weather_id))
	err = obj.db.Table("user_history").Where("user_id = ? and weather_id = ?", id, weather_id).Limit(30).First(&history).Error
	return history, err
}

func (obj *userRepository) Readed(id string) (history []entities.UserHistory, err error) {
	logs.Info(fmt.Sprintf("Function is Called: Readed"))
	logs.Debug(fmt.Sprintf("Data: %v", id))
	err = obj.db.Table("user_history").
		Select("user_id, weather_id, weather_name, quantity").
		Where("user_id = ?", id).
		Group("user_id, weather_id, weather_name, quantity").
		Find(&history).Error
	return history, err
}

func (obj *userRepository) Update(id string, user *entities.User) (userResponse *entities.User, err error) {
	logs.Info(fmt.Sprintf("Function is Called: Update"))
	logs.Debug(fmt.Sprintf("Data: %v, %v", id, user))
	err = obj.db.Table("user").Where("id = ?", id).Updates(user).Error
	err = obj.db.Table("user").Where("id = ?", id).First(&user).Error
	return
}

func (obj *userRepository) Delete(id string) error {
	logs.Info(fmt.Sprintf("Function is Called: Delete"))
	logs.Debug(fmt.Sprintf("Data: %v", id))
	return obj.db.Table("user").Where("id=?", id).Delete(&entities.User{}).Error
}
