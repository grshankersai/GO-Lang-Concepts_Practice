package models

import (
	"errors"

	"shanker.com/rest-service/db"
	"shanker.com/rest-service/utils"
)

type User struct{
	ID int64
	Email string `binding:"required"`
	Password string `binding:"required"`
}


func (u *User) Save() error{
	query := "INSERT INTO users(email,password) VALUES (?,?)"
	stmt , err := db.DB.Prepare(query)

	if err != nil{
		return err
	}

	defer stmt.Close()

	hashedPassword , err :=  utils.HashPassword(u.Password)
	if err != nil{
		return err
	}

	result , err := stmt.Exec(u.Email, hashedPassword)

	if err != nil{
		return err
	}

	userId, err := result.LastInsertId()

	u.ID = userId

	return err
}


func (user *User) ValidateCredentials() error{
	query := "SELECT id,password FROM users WHERE email=?"
	row := db.DB.QueryRow(query,user.Email)

	var retrivedPassword string
	err := row.Scan(&user.ID,&retrivedPassword)

	if err != nil{
		return errors.New("Credentials invalid")
	}

	 passwordIsValid := utils.CheckPasswordHash(user.Password,retrivedPassword)

	 if !passwordIsValid{
		return errors.New("Credentials invalid")
	 }

	 return nil

}