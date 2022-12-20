package main

import "fmt"

//constant values in golang
const Pascal_var = 3000 // public availability

func main() {
	var user string = "Krishna"
	fmt.Println(user)
	fmt.Printf("type of user is: %T \n", user)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("type of user is: %T \n", smallVal)

	var small_float float32 = 255.83
	fmt.Println(small_float)
	fmt.Printf("type of user is: %T \n", small_float)

	//defaultvalues in golang
	var newVar int
	fmt.Println(newVar)
	fmt.Printf("type of user is: %T \n", newVar)

	//implicit types in golang
	var implicit_var = "krishna"
	fmt.Println(implicit_var)
	fmt.Printf("type of user is: %T \n", implicit_var)

	// Walrus operator in golang i.e no variable keyword "var" and variable type needed. only initialisation is required
	walrus_var := "krishna"
	fmt.Println(walrus_var)
	fmt.Printf("type of user is: %T \n", walrus_var)

}
