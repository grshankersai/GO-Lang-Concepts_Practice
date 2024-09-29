package list

import "fmt"

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.



func main(){

	// 1
	hobbies := [3]string{"Eating","Sleeping","Repeat"}
	fmt.Println(hobbies)

	fmt.Println()

	// 2
	fmt.Println("First Element: ", hobbies[0])
	fmt.Println("Second and Third Element: ", hobbies[1:])

	fmt.Println()

	// 3
	// hobbiesSlice := hobbies[:2]
	hobbiesSlice := hobbies[0:2]
	fmt.Println(hobbiesSlice)

	// 4
	latestHobbiesSlice := hobbiesSlice[1:3]
	fmt.Println(latestHobbiesSlice)

	// 5
	goals := []string{"Excel","Empower"}
	fmt.Println(goals)

	// 6
	goals[1] = "Earn"
	goals = append(goals,"Relax")
	fmt.Println(goals)

	// 7
	type Product struct{
		title string
		id string 
		price float64
	}

	productList := []Product{ {title: "Book",id: "P102",price: 200} ,{title: "Pen",id: "P103",price: 100}}
	fmt.Println(productList)

	productList = append(productList, Product{title: "Cycle",id: "P104",price: 2000})

	fmt.Println(productList)

	// IMP syntax destructureing

	prices := []float64{10.99, 9.99, 45.99, 20.0}

	discountPrices := []float64{22.1,20.3,10.5}

	finalSlice := append(prices[:] , discountPrices...)
	fmt.Println(finalSlice)





	
}
















// func main(){
// 	prices := []float64{10.99 , 8.99}

// 	fmt.Println(prices[1])

// 	// Appending an element to an array

// 	newPrices := append(prices,12.69)


// 	// prices := append(prices,12.69)

// 	fmt.Println(newPrices,prices)

// 	// Removing an element
// 	// newPrices = newPrices[1:]


// }














// func main(){

// 	// Method - 1 to create an array.

// 	prices := [4]float64{10.99, 9.99, 45.99, 20.0}

// 	fmt.Println(prices)

// 	// Method - 2
// 	var productNames [3]string  // Declare
// 	productNames = [3]string{"A Book"} // initialization.

// 	fmt.Println(productNames)

// 	// Target Individual values in an array..

// 	fmt.Println(prices[2])

// 	productNames[2] = "A Carpet"
// 	fmt.Println(productNames)

// 	// Slice Feature - 

// 	featuredPrices := prices[1:]
// 	highlighedPrices := featuredPrices[:1]
// 	// featuredPrices := prices[:3] // start from beg uptil lastindx-1
// 	// featuredPrices := prices[1:] // start from 1 till end.

// 	fmt.Println("SLiced Array of prices - ",featuredPrices)


// 	// modify in slice -> modified in main array as well.
// 	// length and capacity
// 	// always can select more elements to the right but not to the left.

// 	fmt.Println("FeaturedPrices",len(featuredPrices),cap(featuredPrices))
// 	fmt.Println("highlighedPrices",len(highlighedPrices),cap(highlighedPrices))


// 	// ------------



// }