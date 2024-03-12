package main

import "fmt"

func main()  {
	fmt.Println("pointers")

	// var ptr *int
	// fmt.Println("ptr is nil:", ptr)
	
	num := 23

	var ptr = &num
	fmt.Println("value of actual pointer is ", ptr)
	fmt.Println("value of actual pointer is ", *ptr)

	*ptr = *ptr*2
	fmt.Println("new value is ", num)
	 
}