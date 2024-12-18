package main

import (
	"fmt"
	"strconv"
	"time"
)

func doSomeTask(str string) {
	fmt.Printf("start task %s\n", str) // Use Printf for formatted output
	time.Sleep(1 * time.Second)
	fmt.Printf("end task %s\n", str)
}

func main() {
	fmt.Println("main 1")
	for i := 1; i < 5; i++ {
		go doSomeTask(strconv.Itoa(i))
	}
	fmt.Println("main 2")
	time.Sleep(5 * time.Second)
}
