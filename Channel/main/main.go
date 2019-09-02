package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Ciao")
	naturals := make(chan int) //si possono definire come delle code che poossono essere bloccanti
	squares := make(chan int)
	//counter thread
	go (func() {
		for i := 1; ; i++ {
			naturals <- i // invio al channell il valore
			time.Sleep(time.Second)
		}
	})()

	//squares
	go (func() {
		for {
			x := <-naturals // consegno ad x il valore appena pushato nel canale naturals
			squares <- x * x
			time.Sleep(time.Second)
		}
	})()

	for {
		//uso il for per bloccare l'uscira perché printLn si mette in attesa
		//quando leggo dal channel esso si svuotas
		fmt.Println(<-squares)
	}
}
