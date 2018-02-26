package main

import (
	"fmt"
)

func main() {
	for i := 65; i < 90; i++ {
		//fmt.Printf("%s:%d\n", i, i)
		fmt.Printf("%c:%d\n", i, i)
	}
	for i := 97; i < 123; i++ {
		//fmt.Printf("%s:%d\n", i, i)
		fmt.Printf("%c:%d\n", i, i)
	}
	// Вывод символа
	fmt.Println(string(65), "\n")
}
