package main

import "fmt"

func main()  {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["rb"] = "ruby"
	languages["py"] = "python"

	fmt.Println("Languages map: ", languages)
	fmt.Println("js stands for:", languages["js"])
	delete(languages, "rb")
	fmt.Println("After deleting from the map: ", languages)

	// loops in golang for maps

	for key,value := range languages{
		fmt.Printf("For key %v, value is %v\n",key,value) 	//%v for value
	}

	for _,value := range languages{
		fmt.Printf("value is %v\n",value)
	}
}