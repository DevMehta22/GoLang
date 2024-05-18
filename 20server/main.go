package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main()  {
	fmt.Println("Web verb")
	GetRequest()
	PostRequest()
	GetRequest()
	//PostFormRequest()
}

func GetRequest()  {
	const myurl = "http://localhost:3000/api/todo/getTodos"

	response,err := http.Get(myurl)
	if err !=nil{
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length:",response.ContentLength)


	var responseString strings.Builder
	content,_ := io.ReadAll(response.Body)
	byteCount,_ := responseString.Write(content)

	fmt.Println("Byte Count is:",byteCount)
	// fmt.Println("Response String:\n"+responseString.String())

	fmt.Println(string(content))

}

func PostRequest()  {
	
	const myurl = "http://localhost:3000/api/todo/addTodo"

	requestBody := strings.NewReader(`
		{
			"title":"GoLang's First post request"
		}
	`)
	
	response,err := http.Post(myurl,"application/json",requestBody)

	if err != nil{
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("Status code:", response.StatusCode)

	content,_ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PostFormRequest(){
	const myurl = "http://localhost:3000/api/todo/addTodo"

	//formdata

	data := url.Values{}

	data.Add("firstname","Dev")
	data.Add("Lastname","Mehta")
	data.Add("email","mehtadev92@gmail.com")

	response,err := http.PostForm(myurl,data)
	if err != nil{
		panic(err)
		}
	defer response.Body.Close();
	fmt.Println("Status code:", response.StatusCode)
	content,_ := io.ReadAll(response.Body)	
	fmt.Println(string(content))
}