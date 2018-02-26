package main

import (
	"fmt"
)

/*
func main() {
		stop := time.After(3 * time.Second)
		tick := time.NewTicker(1 * time.Second)
		defer tick.Stop()
		for {
			select {
			case <-tick.C:
				fmt.Println(time.Now())
			case <-stop:
				return
			}
		}*/
/*
	const filename = "/tmp/file.txt"

	err := ioutil.WriteFile(filename, []byte("Hello, file system\n"), 0644)
	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", b)
}
*/

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
