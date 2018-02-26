package main

// Тур по Go
// Еще типы: структуры, срезы и карты (словари).
// Упражнение: замыкание Фибоначчи 26/27
import "fmt"

// Вывод числа Фибоначчи
func main() {
	var x int = 11

	fibonaci(x)
}

func fibonaci(n int) {
	fmt.Printf("Числа фибоначчи %v ", 0)
	a, b := 0, 1
	for i := 0; i < n-1; i++ {
		a, b = b, a+b
		fmt.Printf("%v ", a)
	}
}
