package main

import (
	"fmt"
	"time"
)

type UserTest struct {
	UserId int `json:"user_id,string" gorm:"type:bigint(20) unsigned auto_increment;not null;primary_key"`
	Nickname string `json:"nickname" gorm:"type:varchar(50);not null;default:'';comment:'昵称'"`
	Avatar string ` json:"avatar" gorm:"type:varchar(35);not null;default:'';comment:'头像'"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" gorm:"column:deleted_at;null"`
}

func NewUserTest() *UserTest {
	return new(UserTest)
}

func AddUser() (bool, error) {
	user := NewUserTest()
	user.Nickname = "A"
	user.Avatar = "B"

	err := DB.Create(user).Error
	if err != nil {
		//fmt.Println(err)
		return false, err
	}

	return true, nil
}

func UpdateUserNickname(userId int, nickname string) (bool, error) {
	update := make(map[string]interface{})
	update["nickname"] = nickname

	fmt.Printf("nickname update: %s .....", nickname)

	err := DB.Where("id = ?", userId).UpdateColumns(update).Error
	if err != nil {
		return false, nil
	}

	return true, nil
}

func AddUserCache() (bool, error) {
	_, err := RedisClient.Set("a", "b", 0).Result()
	if err != nil {
		//fmt.Println(err)
		return false, err
	}

	return true, nil
}
