package main

import "fmt"

func main() {
	// Pointer
	x := 5
	xAddress := &x
	fmt.Println("xAddress =", xAddress)
	fmt.Println("x =", *xAddress)
	*xAddress = 12
	fmt.Println("after x =", x)

	// Array
	arr := [3]int{0, 1000000000000000000, 2000}
	fmt.Println("arr[2] =", arr[2])
	fmt.Println("arr =", arr)

	// Slice
	slice := []string{"This", "เขม", "ja"}
	fmt.Println("addressSlice[0] =", &slice[0])
	fmt.Println("addressSlice[1] =", &slice[1])
	fmt.Println("addressSlice[2] =", &slice[2])
}
