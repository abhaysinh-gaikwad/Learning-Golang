package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [5] string 
	fruitList[0] = "Apple"
	fruitList[1] = "Orange"
	fruitList[4] = "Banana"

	fmt.Println("list of fruits is", fruitList)
	fmt.Println("length of array is", len(fruitList))

	var vegList = [3] string {"potato", "beans", "tomato"}
	fmt.Println("list of vegetables is", vegList)
	fmt.Println("length of array is", len(vegList))
}