package models

import (
	"errors"
	"gopkg.in/mgo.v2/bson"
)

var (
	UserList map[string]*User
)

func init() {
	UserList = make(map[string]*User)
	u := User{"user_11111", "1111", "jeetbangoria","jeetbangoria@gmail.com"}
	UserList["user_11111"] = &u
}

type User struct {
	ID       bson.ObjectId `json:"id" bson:"_id"`
	Password string `json:"password" bson:"password"`
	Username     string `json:"username" bson:"username"`
	Email    string  `json:"email" bson:"email"`
}

func GetUser(uid string) (u *User, err error) {
	if u, ok := UserList[uid]; ok {
		return u, nil
	}
	return nil, errors.New("User not exists")
}

func GetAllUsers() map[string]*User {
	return UserList
}

func UpdateUser(uid string, uu *User) (a *User, err error) {
	if u, ok := UserList[uid]; ok {
		if uu.Username != "" {
			u.Username = uu.Username
		}
		return u, nil
	}
	return nil, errors.New("User Not Exist")
}

func Login(username, password string) bool {
	for _, u := range UserList {
		if u.Username == username && u.Password == password {
			return true
		}
	}
	return false
}

func DeleteUser(uid string) {
	delete(UserList, uid)
}
