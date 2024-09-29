package main

import (
	"fmt"
	"math"
)

func main(){
	fmt.Println("Investment Calculator");

	const inflationRate float64= 2.5

	// var investmentAmount float64 = 1000
	// var expectedReturnRate = 5.5
	expectedReturnRate := 5.5
	// var years float64= 10

	// var investmentAmount  , years  float64= 1000 , 10

	investmentAmount  , years  := 1000.0 , 10.0

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Println()

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	fmt.Println()

	fmt.Print("Enter the expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)


	// var sample int

	// fmt.Scan(&sample)	

	// var futureValue = float64(investmentAmount) * math.Pow((1 + expectedReturnRate / 100),float64(years))
	var futureValue = investmentAmount * math.Pow((1 + expectedReturnRate / 100),years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100,years)


	fmt.Println(futureValue)
	fmt.Println(futureRealValue)


}