package main

import (
	"bufio"
	"os"
)

func main() {
	f, err := os.Create("file.txt")
	if err != nil {
		panic(err)
	}
	size, err := f.WriteString("Hello, World!")
	if err != nil {
		panic(err)
	}
	println(size)
	f.Close()

	archive, _ := os.Open("file.txt")
	reader := bufio.NewReader(archive)
	buffer := make([]byte, 100)
	for {
		_, err := reader.Read(buffer)
		if err != nil {
			break
		}
		println(string(buffer))
	}
}
