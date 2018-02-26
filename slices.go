package main

// Тур по Go
// Еще типы: структуры, срезы и карты (словари).
// Упражнение: срезы 18/27
import "fmt"

func main() {
	var dx = 3
	var dy = 3

	var slice = make([][]uint8, dy)

	fmt.Println(slice)

	for x := 0; x < dx; x++ {
		for y := 0; y < dx; y++ {
			// v := uint8(x * y)
			// v := uint8((x + y) / 2)
			v := uint8(x ^ y)
			slice[x] = append(slice[x], v)
		}
	}
	fmt.Println(slice)

	for i := 0; i < len(slice); i++ {
		for j := 0; j < len(slice); j++ {
			if j == 0 {
				fmt.Printf(" \n")
			}
			fmt.Printf(" %d ", slice[i][j])
		}
	}

}
