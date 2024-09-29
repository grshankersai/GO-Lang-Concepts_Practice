package main

import (
	"fmt"

	"shanker.com/capstone/filemanager"
	"shanker.com/capstone/prices"
)



func main(){	
	var taxRates []float64 = []float64{0,0.07,0.1,0.15}	
	// doneChans := make([]chan bool , len(taxRates))

	for _,taxRate := range taxRates{
		// doneChans[index] =  make(chan bool)
		fm := filemanager.New("prices.txt",fmt.Sprintf("result_%.0f.json",taxRate*100))
		// cmdm := cmdmanager.New()
		priceJob:=prices.NewTaxIncludedPriceJob(fm,taxRate)
		err:=priceJob.Process()
		
		if err != nil{
			fmt.Println("Couldnt process job")
			fmt.Println(err)
		}
	}

	

}