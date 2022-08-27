package main

import (
	"fmt"
)

var z int
var y = 43

func main() {
	fmt.Println("Hello Micah!")
	foo()
	fmt.Println("Print somemore")
	x := 42
	fmt.Println(x)
	x = 99
	fmt.Println(x)

	// for i := 0; i < 100; i++ {
	// 	if i%2 == 0 {
	// 		fmt.Println(i)
	// 	}
	// }
	fmt.Println(z)
	fmt.Println(y)

}

func foo() {
	fmt.Println("I'm in foo, where are you?")
}
