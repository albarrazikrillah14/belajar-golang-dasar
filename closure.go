package main

import "fmt"

func main() {
	counter := 0

	increment := func() {
		fmt.Println("increment")
		counter += 1
	}

	increment()
	increment()
	increment()

	fmt.Println(counter)
}
