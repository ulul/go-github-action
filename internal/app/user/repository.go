package user

import "gorm.io/gorm"

type Repository interface {
	Get() ([]User, error)
	FindByID(id string) (User, error)
	Create(user User) (User, error)
	Delete(id string) error
	UpdateByID(id string, user User) (User, error)
	Login(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Get() ([]User, error) {
	var users []User

	if e := r.db.Find(&users).Error; e != nil {
		return nil, e
	}
	return users, nil
}

func (r *repository) FindByID(id string) (User, error) {
	var user User
	if e := r.db.Where("id = ?", id).First(&user).Error; e != nil {
		return user, e
	}
	return user, nil
}

func (r *repository) Create(user User) (User, error) {
	if e := r.db.Create(&user).Error; e != nil {
		return user, e
	}
	return user, nil
}

func (r *repository) Delete(id string) error {
	var user User
	if e := r.db.Where("id = ?", id).Delete(&user).Error; e != nil {
		return e
	}
	return nil
}

func (r *repository) UpdateByID(id string, user User) (User, error) {
	if e := r.db.Where("id = ?", id).Updates(&user).Error; e != nil {
		return user, e
	}
	return user, nil
}

func (r *repository) Login(user User) (User, error) {
	if e := r.db.Where("username = ? AND password = ?", user.Username, user.Password).First(&user).Error; e != nil {
		return user, e
	}
	return user, nil
}
