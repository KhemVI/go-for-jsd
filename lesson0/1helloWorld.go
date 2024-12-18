package main

import "fmt"

func main() {
	var hello string
	hello = "Hello"
	world := "World"
	var helloWorld string = "Hello World"

	fmt.Println(hello + ", " + world + "!")
	fmt.Printf("%s, %s!", hello, world)
	fmt.Printf("\n\n\n")
	fmt.Println(helloWorld)
}
