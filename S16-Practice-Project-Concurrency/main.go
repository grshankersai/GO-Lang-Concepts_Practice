package main

import (
	"fmt"

	"shanker.com/capstone/filemanager"
	"shanker.com/capstone/prices"
)



func main(){	
	var taxRates []float64 = []float64{0,0.07,0.1,0.15}	
	doneChans := make([]chan bool , len(taxRates))
	errorChans := make([]chan error , len(taxRates))

	for index,taxRate := range taxRates{
		doneChans[index] =  make(chan bool)
		errorChans[index] =  make(chan error)
		fm := filemanager.New("prices.txt",fmt.Sprintf("result_%.0f.json",taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob:=prices.NewTaxIncludedPriceJob(fm,taxRate)
		go priceJob.Process(doneChans[index],errorChans[index])
		// priceJob.Process()
		
		// if err != nil{
		// 	fmt.Println("Couldnt process job")
		// 	fmt.Println(err)
		// }
	}

	for index := range taxRates{
		select{
		case err := <- errorChans[index]:
			if err != nil{
			fmt.Println("Couldnt process job")
			fmt.Println(err)
		}
		case <- doneChans[index]:
			fmt.Print("DOne")

		}
	}

	// for _ , done := range doneChans{
	// 	<- done
	// }

	

}