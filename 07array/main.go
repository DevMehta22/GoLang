package main

import "fmt"

func main()  {
	fmt.Println("Arrays")

	var fruitlist [4]string

	fruitlist[0] = "Apple"
	fruitlist[1] = "Banana"
	fruitlist[3] = "peach"

	fmt.Println("Fruitlist is:", fruitlist)
	fmt.Println(len(fruitlist))

	var veglist = [5]string{"potato","brinjal","beans"}
	fmt.Println("Veglist is : ", veglist, len(veglist))
	
}