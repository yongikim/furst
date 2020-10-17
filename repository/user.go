package repository

import (
	"furst/model"
)

type UserRepository struct{}

func (UserRepository) SetUser(user *model.User) error {
	result := db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (UserRepository) GetUserList() []model.User {
	users := make([]model.User, 0)
	result := db.Limit(10).Find(&users)
	if result.Error != nil {
		panic(result.Error)
	}
	return users
}

func (UserRepository) UpdateUser(newUser *model.User) error {
	result := db.Save(&newUser)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (UserRepository) DeleteUser(id int) error {
	result := db.Delete(&model.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (UserRepository) FindByID(id uint) (model.User, bool) {
	user := model.User{}
	result := db.First(&user, id)
	if result.Error != nil {
		return user, false
	}
	return user, true
}
