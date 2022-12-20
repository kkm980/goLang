package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// walrus constant
	welcome := "Welcome to the user input"
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(welcome)
	// fmt.Printf("type of user is: %T \n", welcome)
	// reader := bufio.NewReader(os.Stdin)
	fmt.Printf("enter the rating for pizza ")

	input, _ := reader.ReadString('\n')
	fmt.Println("thanks for rating ", input)
	fmt.Printf("type of user is: %T \n", input)
}
