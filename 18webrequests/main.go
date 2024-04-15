package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://courses.learncodeonline.in/"

func main()  {
	fmt.Println("LCO web request")

	response,err := http.Get(url)

	if err!=nil {
		panic(err)
	}

	fmt.Printf("Response is of type %T\n",response)

	defer response.Body.Close()

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content :=string(databytes)
	fmt.Println(content)
}