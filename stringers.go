package main

// Тур по Go
// Методы и интерфейсы.
// Упражнение: Stringers 18/26
import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
	fmt.Println(IPAddr{1, 2, 3, 44})
}

func (i IPAddr) String() string {
	return "Хер Вам!!"
	//return strconv.Itoa(int(i[0])) + " - " + strconv.Itoa(int(i[1]))
	// Возврат в конвектруемом формате
}
