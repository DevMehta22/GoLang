package main

import (
	"fmt"
	"math/rand"
	
)

func main()  {
	fmt.Println("Switch Case in golang")

	diceNum := rand.Intn(6) + 1
	fmt.Println("value of dice is:",  diceNum) // roll the die and print its value
	// switch case statement
	switch diceNum {
	case 1:
		fmt.Println("You can open")
	case 2:
		fmt.Println("you can move two spots if already opened")
	case 3:
		fmt.Println("you can move three spots if already opened")
	case 4:
		fmt.Println("you can move four spots if already opened")
	case 5:
		fmt.Println("you can move five spots if already opened")
	case 6:
		fmt.Println("you can move six spots if already opened and have one extra turn")
	default:
		fmt.Println("wrong Input!")
	}

	//fallthrough: Runs the case which is beneath  it, too.
}