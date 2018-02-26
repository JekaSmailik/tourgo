package main

// Тур по Go
// Еще типы: структуры, срезы и карты (словари).
// Упражнение: карты 23/27
import (
	"fmt"
	"strings"
)

func main() {
	x := make(map[string]int)

	var s string = "I ate a donut. Then I ate another donut."

	fmt.Println("Длина 'strings.Fields(s)' =", len(strings.Fields(s)))

	for _, value := range strings.Fields(s) {
		x[value] = 0
		for _, odin := range strings.Fields(s) {
			if value == odin {
				x[value] += 1
			}
		}
	}

	fmt.Println(x)

	fmt.Println("\n ")
	fmt.Println("Map х =", x)
}
