package main

import "fmt"

func main() {
	tab := func() {
		fmt.Println("This is an anonymous function")
	}
	tab() //invoking the function

	func() {
		fmt.Println("This is also anonymous function")
	}()
}
