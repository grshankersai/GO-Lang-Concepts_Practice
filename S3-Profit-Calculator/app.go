package main

import "fmt"

func main(){

	fmt.Println("-------- Profit Calculator-----------")

	var revenue,expense,tax float64

	fmt.Print("Enter the revenue: ")
	fmt.Scan(&revenue)

	fmt.Println()

	fmt.Print("Enter the expense: ")
	fmt.Scan(&expense)

	fmt.Println()

	fmt.Print("Enter the tax: ")
	fmt.Scan(&tax)

	ebt := revenue - expense

	profit := (1- tax/100) * ebt

	ratio := ebt/profit

	fmt.Printf("Ebbitta %v \n",ebt)

	fmt.Println("Ebbita ",ebt)
	fmt.Println("profit ",profit)
	fmt.Println("ratio ",ratio)


	sampleVal := 1234.88791

	fmt.Printf("Sample value print %.2f",sampleVal)


	// Get Strings - 

	ebbitaString := fmt.Sprintf("Ebbitta %v \n",ebt)

	fmt.Print(ebbitaString)

	// USing multi line string:

	multiLine := fmt.Sprintf(`The estimated ebitta value is - 
	
	Ebitta %v`,ebt);

	fmt.Print(multiLine)



}