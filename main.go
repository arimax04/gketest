package main

import (
	"log"
	"runtime"
)

func main() {
	log.Println("Hello, World")
	log.Printf("cpu: %d\n", runtime.NumCPU())
	log.Printf("GOMAXPROC: %d\n",runtime.GOMAXPROCS(0))
}
