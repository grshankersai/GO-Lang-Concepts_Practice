package prices

import (
	"fmt"

	"shanker.com/capstone/conversion"
	"shanker.com/capstone/iomanager"
)



type TaxIncludedPriceJob struct{
	IOManager iomanager.IOManager `json:"-"`
	// IOManager cmdmanager.CMDManager `json:"-"`
	TaxRate float64 `json:"tax_rate"`
	InputPrices []float64 `json:"input_prices"`
	TaxIncludedPrices map[string]string `json:"tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() error{
	lines,err:= job.IOManager.ReadLines()
	if err != nil{
		// fmt.Println(err)
		return err
	}

	prices,err := conversion.StringsToFloat(lines)
	if err != nil{
			// fmt.Println(err)
			return err
	}

	job.InputPrices = prices
	return nil;

}


func (job *TaxIncludedPriceJob) Process() error{
	err :=job.LoadData()

	if(err != nil){
		return err
	}

	result := make(map[string]string)
	for _ , price := range job.InputPrices{
			taxIncludedPrice := price + (price*job.TaxRate)
			result[fmt.Sprintf("%.2f",price)] = fmt.Sprintf("%.2f",taxIncludedPrice)
		}
	job.TaxIncludedPrices = result

	// filemanager.WriteJSON(fmt.Sprintf("./result_%.0f.json",job.TaxRate*100),job)
	return job.IOManager.WriteResult(job)

	
}

func NewTaxIncludedPriceJob( fm iomanager.IOManager,taxRate float64) *TaxIncludedPriceJob{
	return &TaxIncludedPriceJob{
		IOManager: fm,
		InputPrices: []float64{10,20,30},
		TaxRate: taxRate,
	}
}