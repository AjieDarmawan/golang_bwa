package user

import (
	"gorm.io/gorm"
)

// func dbConn() (db *sql.DB) {
// 	dbDriver := "mysql"
// 	dbUser := "root"
// 	dbPass := ""
// 	dbName := "bwastartup"
// 	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return db
// }

// var tmpl = template.Must(template.ParseGlob("form/*"))

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindAll() ([]User, error)
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
	err := r.db.Raw("SELECT id, name, Email FROM users WHERE id = ?", 2).Scan(&res)
	if err != nil {
		return res, nil
	}

	return res, nil
}
