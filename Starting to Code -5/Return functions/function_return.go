package main

import "fmt"

func main() {
	total := addition(2, 5)
	fmt.Println(total)
	div := division(25, 5)
	fmt.Println(div)
}
func addition(a, b int) int {
	total := a + b
	return total //returning addition function
}
func division(c, d float64) float64 {
	div := c / d
	return div //returning division function
}
