package main

import (
	"fmt"
	"os"
)

func main() {
	/*
		i := 0 implicita var e type nella def di variabile(var i int = 0  )
	*/
	for i := 0; i < len(os.Args); i++ {
		fmt.Println(os.Args[i])
	}

	//foreach ( usiamo il blank identify _ per ignorare la chiave altrimenti non compila perchè non la usiamo )
	for _, arg := range os.Args {
		fmt.Println(arg)
	}

	//i:= 0
	// while
	//for i<100{
	//	body
	//}

	// for infinito
	// for{
	//	body
	// }
}
