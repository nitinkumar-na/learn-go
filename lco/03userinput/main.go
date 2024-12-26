package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	fmt.Println("Enter the rating for our Pizza:")
	reader := bufio.NewReader(os.Stdin)

	// comma ok || err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating:", input)
	fmt.Printf("Type of this rating is %T", input)
}