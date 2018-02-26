package main

// Тур по Go
// Методы и интерфейсы.
// Упражнение: Errors 20/26
import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (x ErrNegativeSqrt) Error() string {
	return "cannot MySqrt negative number: " + fmt.Sprint(float64(x))
}

func main() {
	var b ErrNegativeSqrt = 16
	var b2 ErrNegativeSqrt = -16

	fmt.Println(b.Sqrt())
	fmt.Println(b2.Sqrt())
}

func (x ErrNegativeSqrt) Sqrt() (float64, error) {

	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	var niuton float64 = 1
	var srav float64

	for n := 1; n < 30; n++ {
		niuton = niuton - (niuton*niuton-float64(x))/(2*niuton)
		if math.Abs(srav-niuton) <= 0.0001 {
			break
		}
		srav = niuton
	}

	return niuton, nil
}
