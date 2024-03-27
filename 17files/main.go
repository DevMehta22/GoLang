package main

import (
	"fmt"
	"io"
	"os"
)

func main()  {
	fmt.Println("Files in GoLang")
	content := "GoLang programming language"

	file,err := os.Create("./myGoFile.txt")

	if err != nil{
		panic(err)
	}
	
	length,err := io.WriteString(file, content)

	if err != nil{
		panic(err)
	}
	fmt.Printf("The number of bytes written is %d\n", length)

	// Close the file after writing to it
	defer file.Close()

	readFile("./myGoFile.txt")

}

func readFile(filename string)  {
	databyte,err := os.ReadFile(filename)
	
	if err != nil{
		panic(err)
	}
	fmt.Println(string(databyte))
}