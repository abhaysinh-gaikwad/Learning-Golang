package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["rb"] = "ruby"

	fmt.Println("all languages are", languages)
	fmt.Println("js value is", languages["js"])

	delete(languages, "js")
	fmt.Println("all languages are", languages)

	// loops are intersting in golang
	for key, value := range languages {
		fmt.Printf("for key %value value is %v\n", key, value)
	}
	for _, value := range languages {
		fmt.Printf("for key value value is %v\n", key, value)
	}
}
