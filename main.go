package main

import "fmt"

func main() {
	test := []string{"f", "f", "@", "#", "z", "x"}

	fmt.Println(cap(test))
	test = append(test, "Q")
	fmt.Println(cap(test))

	fmt.Println("\n")
	for _, i := range test {
		fmt.Println(i)
	}
}

/**
* pointer: *array
* length: int
* capacity: int
**/

/**
* Checking slice to the void
* : len(slice) == 0 => true|false
**/
