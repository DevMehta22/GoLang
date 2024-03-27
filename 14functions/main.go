package main

import "fmt"

func main()  {
	fmt.Println("Functions in golang")
	greeter()

	result := adder(3,5)
	fmt.Println(result)

	proresult := proadder(2,3,4,5,6,7,8,9)
	fmt.Println(proresult)
}

func greeter() {
	fmt.Println("Namste from GoLang")
}

func adder(val1 int,val2 int) int {
	return val1+val2
}
func proadder(values ... int) int  {
	total := 0
	for _,val := range values{
		total += val
	}
	return total
}