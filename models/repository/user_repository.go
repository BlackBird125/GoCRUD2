package repository

import (
	"github.com/BlackBird125/GoCRUD2/db"
	"github.com/BlackBird125/GoCRUD2/models"
)

type UserRepository struct{}

func (_ UserRepository) GetAll() ([]models.User) {
	db := db.GetDB()
	users := []models.User{}
	db.Find(&users)
	return users
}

func(_ UserRepository) GetByID(id int) (models.User) {
	db := db.GetDB()
	var user models.User
	db.Where("ID = ?", id).First(&user)
	return user
}

func(_ UserRepository) Create(user models.User) (models.User, error) {
	db := db.GetDB()
    if err := db.Create(&user).Error; err != nil {
        return user, err
    }
	return user, nil
}

func(_ UserRepository) Delete(id int) (models.User, error) {
	db := db.GetDB()
	var user models.User
    if err := db.Where("ID = ?", id).Delete(&user).Error; err != nil {
        return user, err
    }
	return user, nil
}

func(_ UserRepository) Update(id int,user models.User) (models.User, error) {
	db := db.GetDB()
	var target models.User
    if err := db.Where("ID = ?", id).First(&target).Updates(&user).Error; err != nil {
        return user, err
    }
	return user, nil
}