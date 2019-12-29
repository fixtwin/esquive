package main

import "fmt"

func main() {
	fmt.Println("Hello everyone!")
	foo()
	fmt.Println("Somehting more")

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
	x := 42 + 5
	fmt.Println(x)
	y := 99
	fmt.Println(y)

}

func foo() {
	fmt.Println("I'm in foo")
}
