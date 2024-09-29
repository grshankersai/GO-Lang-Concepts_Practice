package main

import (
	"errors"
	"fmt"
	"os"
)

func main(){
	var revenue,expense,tax float64

	revenue , err1 := getUserInput("Enter the revenue: ")
	if(err1 != nil){
		panic("Revenue not entered properly.")
	}
	expense , err2 := getUserInput("Enter the expense: ")
	if(err2 != nil){
		panic("Expense not defined properly.")
	}

	tax,err3 := getUserInput("Enter the tax: ")
	if(err3 != nil){
		panic("Tax value is incorrect.")
	}


	ebt,profit,ratio := getInsights(revenue,expense,tax)

	printInformation(ebt,profit,ratio)

}


func getUserInput(message string) (returnValue float64, err error){
	fmt.Print(message)
	fmt.Scan(&returnValue)
	if returnValue == 0 || returnValue < 0{
		return 0, errors.New("Invalid Input Value")
	}

	fmt.Println()
	return
}

func getInsights(revenue float64,expense float64,tax float64) (ebt float64,profit float64,ratio float64) {

	ebt = revenue - expense
	profit = (1- tax/100) * ebt
	ratio = ebt/profit

	return

}

func printInformation(ebt float64 , profit float64,ratio float64){
	fmt.Println("Ebbita ",ebt)
	fmt.Println("profit ",profit)
	fmt.Println("ratio ",ratio)

	storeInfo := fmt.Sprintf("Ebbita %v \n profit %v \n ratio %v \n",ebt,profit,ratio)

	os.WriteFile("Information.txt",[]byte(storeInfo),0644)
	
}