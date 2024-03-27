package main

import "fmt"

func main()  {
	fmt.Println("loops in GoLang")

	days := []string{"Sunday","Tuesday","wednesday","Thursday","Friday","Saturday"}
	
	// fmt.Println(days)

	//for loop

	// for i:=0; i<len(days); i++{
	// 	fmt.Println(days[i])
	// }

	// for i := range days{
	// 	fmt.Println(days[i])
	// }

	for index,day := range days{
		fmt.Printf("Index is %v and value is %v\n",index, day)
	}

	rogueValue := 1

	for rogueValue < 10{
		// if rogueValue==5 {
		// 	break
		// }
		if rogueValue==6{
			rogueValue++
			continue
		}
		if rogueValue==7{
			goto lco
		}
		fmt.Println("value is: ",rogueValue)
		rogueValue++
	}
	lco:
		fmt.Println("Jumped to lco")
}