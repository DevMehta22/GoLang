package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	// no inheritance in golang; no super or parent

	dev := User{"Dev ", "devmehta@go.dev", true, 19}
	fmt.Println(dev)
	fmt.Printf("Dev details are: %+v\n", dev) // +v prints all fields of the struct
	fmt.Printf("Name is %v and email is %v", dev.UserName, dev.Email)
	dev.GetStatus()
	dev.NewMail()
	fmt.Printf("Name is %v and email is %v", dev.UserName, dev.Email)
}

type User struct {
	UserName string
	Email    string
	Status   bool
	Age      int
}

func (u User) GetStatus() {
	fmt.Println("\nIs User active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "dev@go.dev"
	fmt.Println("New Email has been set to :", u.Email)
} 		//passes a copy and donot manipulate original data in structure 