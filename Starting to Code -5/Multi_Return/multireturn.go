package main

import "fmt"

func main() {
	value, error := division(3, 5)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(value)

}
func division(a, b float64) (float64, error) {
	if b == 0 || b == 0.0 {
		return 0, fmt.Errorf("error") //returning error
	}
	return a / b, nil
}
