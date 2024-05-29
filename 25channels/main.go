package main

import (
	"fmt"
	"sync"
)

func main()  {
	fmt.Println("Channels in golang")

	mych := make(chan int, 2)
	wg := &sync.WaitGroup{}

	// mych <- 5
	// fmt.Println(<-mych)

	wg.Add(2)

	//Recieve ONLY
	go func(ch <-chan int, wg *sync.WaitGroup){
		val,isChannelOpen := <- ch
		// close(mych)			Not Allowed in Recieve only routines
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		wg.Done()
	}(mych,wg)


	//Send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup){
		
		// mych <- 0
		ch <- 6
		close(ch)
		
		wg.Done()
	}(mych,wg)
	
	wg.Wait()
}