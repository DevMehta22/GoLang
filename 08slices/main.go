package main

import (
	"fmt"
	"sort"
)

func main()  {
	fmt.Println("Slices")

	// var fruitlist = []string{"Apple","Mango","Orange"}
	// fmt.Printf("Type of fruitlist is %T", fruitlist)

	// fruitlist = append(fruitlist, "Banana","grapes")
	// fmt.Println("\nAfter appending: ", fruitlist)

	// fruitlist = append(fruitlist[1:4])
	// fmt.Println(fruitlist)

	// fruitlist = append(fruitlist[1:],"hello")
	// fmt.Println(fruitlist)
	
	// fruitlist = append(fruitlist[:3])
	// fmt.Println(fruitlist)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 443
	highScores[2] = 753
	highScores[3] = 643
	// highScores[4] = 923 // gives index out of range error

	highScores = append(highScores, 344,675,876,345) //entire memort allocation happens again

	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	sort.Ints(highScores)
	fmt.Println(highScores)


}