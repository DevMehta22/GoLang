package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main()  {
	fmt.Println("Welcome")
	fmt.Println("Please enter value between  1 and 5: ")

	reader := bufio.NewReader(os.Stdin)
	input,_ := reader.ReadString('\n')

	fmt.Println(input)

	numRating,err := strconv.ParseInt(strings.TrimSpace(input),10,64)

	if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println("Added 1 to value:", numRating + 1)
	}


}