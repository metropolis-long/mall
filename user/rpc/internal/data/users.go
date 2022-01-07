package data

import (
	"errors"
	"fmt"
)

func (u *Users) FindUsers(keys *string) ([]*Users, error) {
	db := Db
	users := []*Users{}
	err := db.Table("users").Select("*").Where("name", &keys).Scan(&users).Error
	if err != nil {
		return nil, errors.New("db found err")
	}
	for e, i := range users {
		fmt.Println(i, e)
	}
	return users, nil
}

func (u *Users) GetUser(ID int64) (*Users, error) {
	user := Users{}
	err := Db.First(&user, ID).Error
	if err != nil {
		fmt.Println(err)
		fmt.Println("//////////////////////////////////////////////////////////")
		return nil, err
	}
	return &user, nil
}
