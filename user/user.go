package main

import "time"

type UserTest struct {
	UserId int `json:"user_id,string" gorm:"type:bigint(20) unsigned auto_increment;not null;primary_key"`
	UniqueId string `json:"unique_id" gorm:"type:varchar(50);not null;default:'';unique_index;comment:'微信openid'"`
	Nickname string `json:"nickname" gorm:"type:varchar(50);not null;default:'';comment:'昵称'"`
	Avatar string ` json:"avatar" gorm:"type:varchar(35);not null;default:'';comment:'头像'"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"column:deleted_at;null"`
}

func NewUserTest() *UserTest {
	return new(UserTest)
}