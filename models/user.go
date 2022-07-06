package models

import (
	"errors"

	"github.com/asaskevich/govalidator"
	"gorm.io/gorm"
)

type User struct {
	GormModel
	Username string `gorm:"not null;unique" json:"username" form:"username" valid:"required~Username is required,length(3|255)~Username must be between 3 and 255 characters"`
	Password string `gorm:"not null" json:"password" form:"password" valid:"required~Password is required,length(6|255)~Password must be between 8 and 255 characters"`
	Email    string `gorm:"not null;unique" json:"email" form:"email" valid:"required~Email is required,email~Email is invalid"`
	Age      int    `gorm:"not null" json:"age" form:"age" valid:"required~Age is required"`
	Name     string `gorm:"not null" json:"name" form:"name" valid:"required~Name is required,length(3|255)~Name must be between 3 and 255 characters"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	_, errCreate := govalidator.ValidateStruct(u)

	if u.Age < 5 {
		err = errors.New("age must be greater than 5")
		return err
	}

	if errCreate != nil {
		err = errCreate
		return
	}

	err = nil
	return
}
