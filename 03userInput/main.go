package main

import (
	"bufio"
	"fmt"
	"os"
)

func main()  {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	// read a line of text from the reader
	fmt.Println("Enter the input:")
	
	// comma ok || err

	input, _ := reader.ReadString('\n')
	fmt.Println(input)
	fmt.Printf("Type of input %T", input)
}