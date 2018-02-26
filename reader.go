package main

// Тур по Go
// Методы и интерфейсы.
// Упражнение: Reader 22/26
import "golang.org/x/tour/reader"

type MyReader struct{}

func (MyReader) Read(b []byte) (int, error) {
	b[0] = 'A'
	return 1, nil
}

// TODO: Add a Read([]byte) (int, error) method to MyReader.

func main() {
	reader.Validate(MyReader{})
}
