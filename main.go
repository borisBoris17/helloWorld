package main

import "fmt"

func main() {
	fmt.Println("Hello Micah!")
	foo()
	fmt.Println("Print somemore")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I'm in foo, where are you?")
}
