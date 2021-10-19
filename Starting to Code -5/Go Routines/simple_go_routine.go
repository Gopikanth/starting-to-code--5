/*package main

import "fmt"

func main() {
	go next("hi")
	next("hello")      //only the normal function is executed
}
func next(str string) {
	for i := 0; i < 5; i++ {
		fmt.Println(str)
	}
} */

package main

import (
	"fmt"
	"time"
)

func main() {
	go next("hi")
	next("hello") //both the normal function  and go routine function is executed
}
func next(str string) {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond)
		fmt.Println(str)
	}
}
