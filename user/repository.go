package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindByID(ID int) (User, error)
	FindAll() ([]User, error)
	Update(user User) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	err := r.db.Create(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(ID int) (User, error) {
	var user User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

// func (r *repository) FindAll() ([]User, error) {
// 	var users []User

// 	err := r.db.Order("id desc").Find(&users).Error

// 	if err != nil {
// 		return users, err
// 	}

// 	return users, nil
// }

func (r *repository) FindAll() ([]User, error) {
	//var users []User
	res := []User{}

	//err := r.db.Raw("SELECT id, name, city FROM users WHERE name = ?", 3).Scan(&res)
	// err := r.db.Raw("SELECT id, name, Email FROM users ").Scan(&res)
	err := r.db.Raw("SELECT * FROM users ").Scan(&res)
	if err != nil {
		return res, nil
	}

	return res, nil
}

func (r *repository) Update(user User) (User, error) {
	err := r.db.Save(&user).Error

	if err != nil {
		return user, err
	}

	return user, nil
}
