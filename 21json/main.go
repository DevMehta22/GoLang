package main

import (
	"encoding/json"
	"fmt"
)

type course struct{
	Name 		string `json:"coursename"`
	Price 		int
	Platform 	string `json:"website"` //alias
	Password 	string `json:"-"`   	//it won't be displayed 
	Tags 		[]string `json:"tags,omitempty"` //won't show null values
}

func main()  {
	fmt.Println("create and handling JSON")
	EncodeJson()


}

func EncodeJson()  {
	
	Courses := []course{
		{"Go Programming", 100, "Udemy", "abc1234", []string{"Programming", "Go"}},
		{"React JS", 200, "Udemy", "lmn5678", []string{"Programming", "React"}},
		{"Python Programming", 300, "Udemy", "xyz9101", nil},
	}

	//package this data as JSON 

	finaljson,err := json.MarshalIndent(Courses,"","\t")
	if err != nil {
		fmt.Println("error in marshalling",err)
		}
	
	fmt.Println(string(finaljson))
}