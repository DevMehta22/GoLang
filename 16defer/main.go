package main

import "fmt"

func main()  {
	fmt.Println("Defers in golang")
	defer fmt.Println("Hello World!") // deferred call to Hello World

	// the function will be executed when control returns from the calling function
	
	defer fmt.Println("Goodbye World!") // another deferred call to Goodbye World
	defer fmt.Println("I'm out of here!") // a third deferred statement
	fmt.Println("Lesgoo.. GoLang!")   // normal execution of code

	//defer functions as Lifo(last in first out)
	mydefer()
}

func mydefer()  {
	for i := 0; i < 5; i++ {
		defer fmt.Println("Loop Value: ", i)
	}
}