package main

// Тур по Go
// Операторы управления потоком: for, if, else, switch и defer.
// Упражнение: циклы и функции 8/14
import (
	"fmt"
	"math"
)

var x float64 = 555

func main() {
	// Запуск кода через три секунды после запуска
	//time.Sleep(3*time.Second)

	fmt.Println("Корень", x, "равен", math.Sqrt(x))

	fmt.Printf("%.5f  Вывод корня по методу Ньютона \n", MySqrt(x))

}

func MySqrt(x float64) float64 {

	var niuton float64 = 1
	var srav float64

	for n := 1; n < 30; n++ {
		niuton = niuton - (niuton*niuton-x)/(2*niuton)
		fmt.Print(n, "комп")
		fmt.Printf("  %.25f \n", niuton)
		//math.Abs() модуль значений "убирает знак минуса"

		if math.Abs(srav-niuton) <= 0.00000001 {
			fmt.Printf("%.11f \n", niuton)
			break
		}
		srav = niuton
	}

	return niuton

}
