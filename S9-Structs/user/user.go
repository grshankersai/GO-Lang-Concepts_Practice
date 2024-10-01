package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct{
	firstName string
	lastName string
	dateOfBirth string
	createdAt time.Time
}

type Admin struct{
	User
	email string
	password string
}

func (u User) OutputUserDetails(){
	fmt.Println(u.firstName,u.lastName,u.dateOfBirth)
}

func(u *User) ClearUserName(){
		u.firstName = ""
		u.lastName = ""
}

func NewUser(firstName,lastName,dateOfBirth string) (*User,error){
	if firstName == "" || lastName == "" || dateOfBirth == ""{
		return nil , errors.New("Invalid data provided")
	}

	return &User{
		firstName: firstName,
		lastName: lastName,
		dateOfBirth: dateOfBirth,
		createdAt: time.Now(),
	},nil
}


func NewAdmin(email,password string) Admin{
	return Admin{
		email: email,
		password: password,
		User: User{
			firstName: "Shanker",
			lastName: "Sai",
			dateOfBirth: "01/05/2000",
			createdAt: time.Now(),

		},
	}

}

