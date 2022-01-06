package main

import (
	"log"
	"time"
)

var logger = log.Default()

func main() {
	count := 0
	for {
		count += 1
		logger.Println("hello")
		time.Sleep(time.Second * 15)
		if count%4 == 0 {
			panic("has error")
		}
	}
}
