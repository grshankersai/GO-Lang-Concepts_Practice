package main

import (
	"fmt"

	"shanker.com/structs/user"
)

// Custome type:

type str string

func (value str) log(){
	fmt.Println(value)
}


func main(){
	firstName := getUserData("Enter your First Name")
	lastName := getUserData("Enter Your Last Name")
	dateOfBirth := getUserData("Enter your DOB")

	//

	// var appUser *user = newUser(firstName,lastName,dateOfBirth)

	// appUser , error := newUser(firstName,lastName,dateOfBirth)

	var appUser  *user.User

	appUser , error := user.NewUser(firstName,lastName,dateOfBirth)

	admin := user.NewAdmin("shanker@email.com","password")

	admin.OutputUserDetails()


	// appUser = user{
	// 	firstName: firstName,
	// 	lastName: lastName,
	// 	dateOfBirth: dateOfBirth,
	// 	createdAt: time.Now(),
	// }

	// assignment alternate way.
	// appUser = user{
	// 	firstName,
	// 	lastName,
	// 	dateOfBirth,
	// 	time.Now(),
	// }

	// outputUserDetails(&appUser)

	if(error != nil){
		fmt.Print("Error")
		return
	
	}

	// fmt.Println(firstName,lastName,dateOfBirth)

	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()


	var val str = "Shanker"

	val.log()




}

// func outputUserDetails(usr *user){
// 	fmt.Print(usr)
// 	fmt.Println(usr.firstName,usr.lastName,usr.dateOfBirth,usr.createdAt)
// }

func getUserData(prompt string) string{
	fmt.Println(prompt)
	var val string
	fmt.Scanln(&val)
	return val
}