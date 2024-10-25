package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	go process(0)

	go func() {
		fmt.Println("its me")
	}()

	for i := 0; i < 1000; i++ {
		go process(i)
	}
}

func process(i int) {
	fmt.Println("settings: ", i)

}
