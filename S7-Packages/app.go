package main

import (
	"fmt"

	"shanker.com/packages/fileops"

	"github.com/Pallinder/go-randomdata"
);

const accountBalanceFile = "balance.txt"



func main(){

	var accountBalance , err = fileops.GetFloatFromFile(accountBalanceFile)

	if(err != nil){
		fmt.Println("ERROR")
		fmt.Println(err)
		// panic("Cant continue")
	}

	fmt.Println("Welcome to GO Bank!!")
	fmt.Println("Reach us 24/7",randomdata.PhoneNumber())
	for {

		presentOptions()

	var choice int
	fmt.Scan(&choice)


	

	if choice == 1 {
		fmt.Println("Your account balance is: ",accountBalance)
	}else if choice == 2 {
		fmt.Println("How much do you want to diposit ?")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		if depositAmount < 0 {
			fmt.Println("Cannot perform negative deposit")
			return
		}else{
			accountBalance += depositAmount
			fileops.WriteFloatToFile(accountBalance,accountBalanceFile)
			fmt.Println("New Account Balance ", accountBalance)
		}
	}else if choice == 3{
		fmt.Print("How much money do you want to withdraw??")
		var withdrawAmount float64
		fmt.Scan(&withdrawAmount)
		if withdrawAmount > accountBalance{
			fmt.Print("Cant perform the operation")
			return
		}else{
			accountBalance -= withdrawAmount
			fileops.WriteFloatToFile(accountBalance,accountBalanceFile)
			fmt.Println("New Account Balance ", accountBalance)
		}
	}else{
		fmt.Print("Good Bye!!!")
		break
	}

	}


}


