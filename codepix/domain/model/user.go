package model

import (
	"errors"
	"strings"
	"time"

	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type User struct {
	Base  `valid:"required"`
	Name  string `json:"name" gorm:"type:varchar(120)" valid:"notnull"`
	Email string `json:"email" gorm:"type:varchar(150)" valid:"notnull"`
}

func (user *User) isValid() error {
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		return err
	}

	if user.Email == "" || !strings.ContainsAny(user.Email, "@.") {
		return errors.New("email not informed or invalid")
	}

	if user.Name == "" || len(strings.Split(user.Name, " ")) <= 1 {
		return errors.New("name is not informed or not completed")
	}

	return nil
}

func NewUser(email, name string) (*User, error) {
	user := User{
		Name:  name,
		Email: email,
	}

	user.ID = uuid.NewV4().String()
	user.CreatedAt = time.Now()

	err := user.isValid()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
