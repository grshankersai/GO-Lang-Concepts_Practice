package main

import "fmt"

func main(){
	// greeting("Shanker")


	// product,sum := calculateMultiplyAndSum(1,2)

	// fmt.Printf("product - %v",product)
	// fmt.Printf("sum - %v",sum)

	var revenue,expense,tax float64

	revenue = getUserInput("Enter the revenue: ")
	expense = getUserInput("Enter the expense: ")
	tax = getUserInput("Enter the tax: ")


	ebt,profit,ratio := getInsights(revenue,expense,tax)

	printInformation(ebt,profit,ratio)



	// ebt := revenue - expense

	// profit := (1- tax/100) * ebt

	// ratio := ebt/profit

	// fmt.Printf("Ebbitta %v \n",ebt)

	// fmt.Println("Ebbita ",ebt)
	// fmt.Println("profit ",profit)
	// fmt.Println("ratio ",ratio)

}


func getUserInput(message string) (returnValue float64){
	fmt.Print(message)
	fmt.Scan(&returnValue)
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
}








func greeting(text string){
	fmt.Printf("Hello %v",text);
}

func calculateMultiplyAndSum(num1 int, num2 int) (sum int,product int){
	sum = num1+num2
	product = num1*num2
	// return num1*num2,num1+num2
	return
}