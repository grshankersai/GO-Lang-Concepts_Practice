package functionsarevalues

import "fmt"


type funcSyntax func(int) int

func main(){

	numbers := []int{1,2,3,4,5}
	// moreNumbers := []int{5,6,7,8}

	newNumbersArray := transformNumbers(&numbers,getTransformerFunction(&numbers))

	fmt.Println(newNumbersArray)

}

func transformNumbers(numbers *[]int, transform funcSyntax) []int{
	var doubleNumbers []int

	for _,value := range *numbers{
		doubleNumbers = append(doubleNumbers,transform(value))
	}
	return doubleNumbers
}


func getTransformerFunction(numbers *[]int) func(int) int{
	if (*numbers)[0] == 1{
		return double
	}else{
		return Triple
	}
}

func double(number int) int{
	return number*2
}

func Triple(number int) int{
	return number*3
}