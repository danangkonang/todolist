package model

import "time"

type Activity struct {
	Id              int       `json:"id" gorm:"primaryKey;column:activity_id"`
	Title           string    `json:"title"`
	Email           string    `json:"email"`
	ActivityGroupId Todo      `json:"-" gorm:"Foreignkey:ActivityGroupId;association_foreignkey:Id;"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
