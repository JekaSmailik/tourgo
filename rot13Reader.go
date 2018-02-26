package main

// Тур по Go
// Методы и интерфейсы.
// Упражнение: roe13Reader 23/26
import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("Spwwd b qgpco!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

func (r rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)
	for i := 0; i < n; i++ {
		b[i] = rot11(b[i])
	}
	return n, err
}
func rot11(b byte) byte {
	if b >= 'a' && b <= 'z' {
		// Поворот строчных букв 13 мест.
		if b >= 'm' {
			return b - 11
		} else {
			return b + 11
		}
	} else if b >= 'A' && b <= 'Z' {
		// Поворот заглавных букв 13 мест.
		if b >= 'M' {
			return b - 11
		} else {
			return b + 11
		}
	}
	// Ничего не делать.
	return b
}

/*
func rot13byte(x byte) byte {
	s := rune(x)
	if s >= 'a' && s <= 'm' || s >= 'A' && s <= 'M' {
		x += 13
	}
	if s >= 'n' && s <= 'z' || s >= 'N' && s <= 'Z' {
		x -= 13
	}
	return x
}*/
