package main

import "fmt"

func main(){
	age := 32

	var agePointer *int

	agePointer = &age

	fmt.Println("age: ",*agePointer)

	getAdultAge(agePointer)

	fmt.Println("after adult age: ",*agePointer)

	// adultAge := getAdultAge(age)
	// fmt.Println(adultAge)


}

func getAdultAge(age *int)  {

	*age = *age - 18

	// return age-18
}