package main

import (
	"fmt"
	"os"
)

func main() {
	i := 0

	for {
		f, err := os.Create(fmt.Sprintf("./tmp/file%d.txt", i))
		if err != nil {
			panic(err)
		}
		defer f.Close()

		_, err = f.WriteString("Hello, World!")
		if err != nil {
			panic(err)
		}

		i++
	}
}
