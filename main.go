package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	for {
		pid := os.Getpid()
		fmt.Printf("Current PID: %d\n", pid)
		time.Sleep(2 * time.Second)
	}
}
