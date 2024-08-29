package main

import "fmt"

 
const pi = 3.14
var jwtToken int = 30000;
 
func main() {
	fmt.Println("this is second class of learing go");

	fmt.Println("value of pi is ", pi)
	var firstname string 
	var lastname string = "kumar"
	var fullname = firstname + " " + lastname
	fmt.Println(fullname)
	fmt.Printf("the type of fullname is %T \n", fullname)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("the type of isLoggedIn is %T \n", isLoggedIn)


	var smallVal uint8 = 255 // 0 to 255 if smallVal > 255 it will give error because it is out of range
	var smallVal2 uint16 = 65535
	var smallVal3 uint32 = 4294967295

	fmt.Println(smallVal)
	fmt.Println(smallVal2)
	fmt.Println(smallVal3)

	fmt.Printf("the type of smallVal is %T \n", smallVal)

	var Value int = 42
	fmt.Println(Value)
	fmt.Printf("the type of Value is %T \n", Value)


	var smallFloat  float64 = 255.58484849784531
	fmt.Println(smallFloat)
	fmt.Printf("the type of smallFloat is %T \n", smallFloat)


	website := "learn go"
	fmt.Println(website)

	website1 := "golang.com" 
	fmt.Println(website1)
}


