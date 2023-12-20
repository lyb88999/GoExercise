package main

import (
	"log"
	"math/rand"
	"time"
)

func sayGreeting(Greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(Greeting)
		interval := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(interval)
	}
}

func main() {
	log.SetFlags(0)
	go sayGreeting("Hello", 100)
	go sayGreeting("Hi", 100)
	time.Sleep(2 * time.Second)
}
