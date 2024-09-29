package maps

import "fmt"

func main(){
	website := map[string]string{
		"Google":"https://www.google.com",
		"AWS":"www.aws.com",
	}

	fmt.Println(website)

	// Extract Information...

	fmt.Println(website["Google"])

	website["apple"]  = "www.apple.com"


	fmt.Println(website)

	delete(website,"apple")

	fmt.Println(website)
}