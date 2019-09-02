package main

import (
	"fmt"

	"github.com/ttacon/chalk" // devi prima importare con go get
)

func main() {
	fmt.Println(
		chalk.Magenta.Color("Hello world ⏰"),
	)
}
