package main

import "fmt"

func main(){
	fmt.Println("Hello world")
	letters := []string{"a", "b", "c"}
	letters =  append(letters, "d")
	fmt.Println(letters)
}