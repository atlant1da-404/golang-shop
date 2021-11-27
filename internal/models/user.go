package models

import "errors"

type User struct {
	Id   int    `sql:"id" json:"id"`
	Name string `sql:"name" json:"name"`
	Age  int    `sql:"age" json:"age"`
}

type UserUpdate struct {
	Model User
}

func (u User) Validation() error {
	if u.Name == "" {
		return errors.New("invalid user")
	}
	if u.Age <= 0 {
		return errors.New("invalid age")
	}
	return nil
}
