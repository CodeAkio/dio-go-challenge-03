package main

import (
	"fmt"
	"time"
)

func ping(pings chan<- string, pongs <-chan string) {
	for {
		<-pongs // Espera receber "pong"
		fmt.Println("ping")
		time.Sleep(time.Second) // Delay para visualização
		pings <- "ping"         // Envia "ping"
	}
}

func pong(pings <-chan string, pongs chan<- string) {
	for {
		<-pings // Espera receber "ping"
		fmt.Println("pong")
		time.Sleep(time.Second) // Delay para visualização
		pongs <- "pong"         // Envia "pong"
	}
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	go ping(pings, pongs)
	go pong(pings, pongs)

	// Inicia o ciclo
	pongs <- "pong"

	// Roda as goroutines por um tempo antes de terminar o programa
	time.Sleep(10 * time.Second)
}
