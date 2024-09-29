package main

import "fmt"

func main(){

	fmt.Println(sumup(1,2,3,4,5))

	numbers := []int{1,2,3}
	fmt.Println(sumup(numbers...))

}

// collect all
func sumup(numbers ...int) int{
	sum := 0;

	for _,val := range numbers{
		sum += val;
	}

	return sum;
}