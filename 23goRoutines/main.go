package main

import (
	"fmt"
	"net/http"
	"sync"
	// "time"
)

var signals = []string{"test"}
var wg sync.WaitGroup //usually these are pointers and not simple variables
var mut sync.Mutex //usually these are pointers and not simple variables

func main()  {
	// go greeter("Hello")
	// greeter("World")

	websitelist := []string{
		"https://www.google.com",
		"https://www.facebook.com",
		"https://www.instagram.com",
		"https://www.youtube.com",
	} 

	for _,web := range websitelist {		
		go getStatusCode(web)
		wg.Add(1)

	}

	wg.Wait()
	fmt.Println(signals)
}

// func greeter(s string)  {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(2*time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string)  {
	defer wg.Done()

	res,err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
		}
	mut.Lock()
	signals = append(signals, endpoint)
	mut.Unlock()

	fmt.Println(res.StatusCode)
}