package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	firstName := "Hunter"
	lastName := "Pittman"

	var d uint8
	d = 255

	fmt.Printf("My name is %v %v \n", firstName, lastName)
	fmt.Printf("My number is %v", d)
	goodbye()

}

func goodbye() {
	fmt.Println("Goodbye!")
}
