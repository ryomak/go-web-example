package model

import "github.com/jinzhu/gorm"

type Question struct {
	gorm.Model
	Text      string
	didAnswer bool
	UserID    uint
	User      *User
}

func GetAllByUserID(db *gorm.DB, userID uint) ([]Question, error) {
	quesions := make([]Question, 10)
	if err := db.Where("user_id = ?", userID).Find(&quesions).Error; err != nil {
		return nil, err
	}
	return quesions, nil
}

func (q *Question) Insert(db *gorm.DB) (*Question, error) {
	if err := db.Create(&q).Error; err != nil {
		return nil, err
	}
	return q, nil
}

func (q *Question) Update(db *gorm.DB) (*Question, error) {
	if err := db.Updates(&q).Error; err != nil {
		return nil, err
	}
	return q, nil
}
