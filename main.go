package main

import (
	"fmt"
	"time"
)

func add(a int, b int) int {
	return a + b
}

func main() {
	sum := add(5, 6)
	fmt.Printf("%v\n", sum)

	for i := 0; i < 10; i++ {
		time.Sleep(time.Minute * 2)
		// Do something every 2min
		const name = "Optum"
		s := fmt.Sprintf("We are %s\n", name)
		fmt.Println(s)
	}
}
