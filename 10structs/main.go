package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; no super or parent

	dev := User{"Dev ","devmehta@go.dev",true,19}
	fmt.Println(dev)
	fmt.Printf("Dev details are: %+v\n", dev) // +v prints all fields of the struct
	fmt.Printf("Name is %v and email is %v", dev.UserName, dev.Email)

}

type User struct {
	UserName string
	Email    string
	Status   bool
	Age      int
}
