package main

import "fmt"

func main() {
	fmt.Println("Peter Jan Singer")

	foo()
	fmt.Println("byeeee..")
}

func foo() {
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}
