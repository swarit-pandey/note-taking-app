package repository

import "github.com/jinzhu/gorm"

type UserRepository interface {
	CreateUser(user User) error
	GetUserByEmail(email string) (User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) error {
	return r.db.Create(&user).Error
}

func (r *userRepository) GetUserByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).First(&user).Error
	return user, err
}
