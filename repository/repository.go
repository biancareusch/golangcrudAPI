package repository

import "time"

//Repository represent the repositories
type Repository interface {
	Close()
	FindByID(id int) (*UserModel, error)
	Find() ([]*UserModel, error)
	Create(user *UserModel) error
	Update(user *UserModel) error
	Delete(id int) error
}

//UserModel represent the user model

type UserModel struct {
	ID int
	FirstName string
	LastName string
	Age int
	DateJoined time.Time
	DateUpdated time.Time
}