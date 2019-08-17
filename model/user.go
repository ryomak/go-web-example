package model

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Name        string `form:"name" binding:"required" gorm:"not null"`
	Mail        string `form:"mail" binding:"required" gorm:"unique;not null"`
	Password    string `form:"password" binding:"required" gorm:"not null"`
	Description string
}

func UserByID(db *gorm.DB, id uint) (*User, error) {
	u := new(User)
	if err := db.Where("id =?", id).First(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func UserByMail(db *gorm.DB, mail string) (*User, error) {
	u := new(User)
	if err := db.Where("mail =?", mail).First(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) Insert(db *gorm.DB, password string) (*User, error) {
	salted, err := passwordHash(password)
	if err != nil {
		return nil, err
	}
	u.Password = salted
	if err := db.Create(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

func (u *User) Update(db *gorm.DB) (*User, error) {
	err := db.Model(&u).Where("id = ?", u.ID).Updates(map[string]interface{}{"name": u.Name, "mail": u.Mail, "Description": u.Description}).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func Auth(db *gorm.DB, mail, password string) (*User, error) {
	user, err := UserByMail(db, mail)
	if err != nil {
		return nil, err
	}
	if err := passwordVerify(user.Password, password); err != nil {
		return nil, err
	}
	return user, nil
}

func passwordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), err
}

func passwordVerify(hash, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
