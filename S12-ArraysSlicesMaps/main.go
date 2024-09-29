package main

import "fmt"

type floatMap map[string]float64;

func (f floatMap) Output(){
	fmt.Println(f)
}

func main(){
	userNames := make([]string, 2,5)

	userNames = append(userNames,"Max")
	userNames = append(userNames,"Manuel")

	fmt.Println(userNames)

	// courseRatings := map[string]float64{}
	courseRatings := make(floatMap,3)

	courseRatings["go"] = 4.7
	courseRatings["React"] = 4.8
	courseRatings["Angular"] = 4.7

	fmt.Println(courseRatings)

	courseRatings.Output()

	// For loop for lists

	for index, value := range userNames{
		fmt.Println(index,value)
	}

	// for loop for maps:

	for key,value := range courseRatings{
		fmt.Println(key,value)
	}



}