package main

import "fmt"

const LoginToken string = "gsdfsdgr" //public

func main() {
	var username string = "Dev"
	fmt.Println(username)
	fmt.Printf("variable is of type %T \n", username)

	var isLoggedin bool = false
	fmt.Println(isLoggedin)
	fmt.Printf("variable is of type %T \n", isLoggedin)

	var smallval int = 255
	fmt.Println(smallval)
	fmt.Printf("variable is of type %T \n", smallval)

	var smallFloat float32 = 255.4354545345
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type %T \n", smallFloat)

	//default values and some aliases
	var variable int
	fmt.Println(variable) //0
	fmt.Printf("variable is of type %T \n", variable)

	//implicit type 
	var website = "devmehta.in"
	fmt.Println(website)

	//no var style (only inside a method not globally)
	numberofuser := 3000
	fmt.Println(numberofuser)

	//public variable
	fmt.Println(LoginToken)
	fmt.Printf("variable is of type %T \n", LoginToken)
}
