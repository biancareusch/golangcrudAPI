package repository

import (
	"time"
)
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
	IDdb int
	FirstNamedb string
	LastNamedb string
	Agedb int
	DateJoineddb time.Time
	DateUpdateddb time.Time
}