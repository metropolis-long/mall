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

func (u *Users) FindUserLike(id int64) (*UserLike,error) {
	query := &Users{}
	user, err := query.GetUser(id)
	if err != nil {
		fmt.Println(err)
		fmt.Println("//////////////////////////////////////////////////////////")
		return nil, err
	}
	likes := []*Like{}
	//.Where("user_id", &id)
	err = Db.Table("likes").Select("*").Scan(&likes).Error
	if err != nil {
		fmt.Println(err)
		fmt.Println("777777777777777777777777777777777777777777777777777")
		return nil, err
	}
	res := &UserLike{}
	res.Likes = likes
	res.User = *user
	return res, nil
}
