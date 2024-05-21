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
	// EncodeJson()
	DecodeJson()


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

func DecodeJson()  {
	jsonData := []byte(`
	{
		"coursename": "Go Programming",
		"Price": 100,
		"website": "Udemy",
		"tags": [
				"Programming",
				"Go"
		]
}
	`)

	var cc course 

	checkValid := json.Valid(jsonData)

	if checkValid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonData, &cc)
		fmt.Printf("%#v\n",cc) 		// # is used to print the type of data and v is used to print the actual value of that data
	}else{
		fmt.Println("JSON is invalid")
	}

	//sone cases were data is to be added as key value pair

	var myOnlineData map[string]interface{}
	
	json.Unmarshal(jsonData,&myOnlineData)
	fmt.Printf("%#v\n",myOnlineData)

	for k,v := range myOnlineData {
		fmt.Printf("key is %v and value is %v and Type is %T\n",k,v,v)
	}

}