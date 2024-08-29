package main

import "fmt"

func main() {
	fmt.Println("learning pointers in golang")
	var ptr *int

	fmt.Println("vale of ptr is", ptr)

	myNumber := 23 
	
	var ptr2 = &myNumber
	fmt.Println("value of number is", ptr2)
	fmt.Println("value of number is", *ptr2)
}