package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
);

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64,error){
	data , err := os.ReadFile(accountBalanceFile)

	if err != nil{
		return 1000, errors.New("Failed to read file")
	}

	balanceText := string(data)
	balanceFloat , err:= strconv.ParseFloat(balanceText,64)

	if err != nil{
		return 1000, errors.New("Failed to parse stored balance value")
	}

	return balanceFloat,nil

}

func writeBalanceToFile(balance float64){
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile,[]byte(balanceText), 0644 )
}

func main(){

	var accountBalance , err = getBalanceFromFile()

	if(err != nil){
		fmt.Println("ERROR")
		fmt.Println(err)
		// panic("Cant continue")
	}

	fmt.Println("Welcome to GO Bank!!")
	for {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Scan(&choice)

	// fmt.Println("Your Choice: ",choice)


	switch choice {
	case 1:
		fmt.Println("Your account balance is: ",accountBalance)
	case 2:
		fmt.Println("How much do you want to diposit ?")
		var depositAmount float64
		fmt.Scan(&depositAmount)
		if depositAmount < 0 {
			fmt.Println("Cannot perform negative deposit")
			return
		}else{
			accountBalance += depositAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("New Account Balance ", accountBalance)
		}
	case 3:
		fmt.Print("How much money do you want to withdraw??")
		var withdrawAmount float64
		if withdrawAmount > accountBalance{
			fmt.Print("Cant perform the operation")
			return
		}else{
			accountBalance -= withdrawAmount
			writeBalanceToFile(accountBalance)
			fmt.Println("New Account Balance ", accountBalance)
		}
	default:
		fmt.Print("Good Bye!!!")
		

	}

	// if choice == 1 {
	// 	fmt.Println("Your account balance is: ",accountBalance)
	// }else if choice == 2 {
	// 	fmt.Println("How much do you want to diposit ?")
	// 	var depositAmount float64
	// 	fmt.Scan(&depositAmount)
	// 	if depositAmount < 0 {
	// 		fmt.Println("Cannot perform negative deposit")
	// 		return
	// 	}else{
	// 		accountBalance += depositAmount
	// 		fmt.Println("New Account Balance ", accountBalance)
	// 	}
	// }else if choice == 3{
	// 	fmt.Print("How much money do you want to withdraw??")
	// 	var withdrawAmount float64
	// 	if withdrawAmount > accountBalance{
	// 		fmt.Print("Cant perform the operation")
	// 		return
	// 	}else{
	// 		accountBalance -= withdrawAmount
	// 		fmt.Println("New Account Balance ", accountBalance)
	// 	}
	// }else{
	// 	fmt.Print("Good Bye!!!")
	// 	break
	// }

	}


}