package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in golang")

	var fruitList = []string{}
	// fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Apple", "Mango", "Orange", "Kiwi")
	// fmt.Println("Fruit list is", fruitList)

	fruitList = append(fruitList[:1])
	// fmt.Println("Fruit list is", fruitList)

	highScore := make([]int, 4) 

	highScore[0] = 234
	highScore[1] = 345
	highScore[2] = 456
	highScore[3] = 867
	
	
	highScore = append(highScore, 555, 666, 321)
	
	// fmt.Println(sort.IntsAreSorted(highScore))
	sort.Ints(highScore)

	// fmt.Println("High score is", highScore[len(highScore)-1:])

	// fmt.Println(sort.IntsAreSorted(highScore))


	var courses = []string{"reactjs", "javascript", "swift", "python", "ruby"}
	fmt.Println("all courses are", courses)

	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println("all courses are", courses)

}
