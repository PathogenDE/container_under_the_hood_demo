package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	for {
		if err := checkAndCreateSecret(); err != nil {
			fmt.Printf("Error creating secret file: %v\n", err)
		}
		pid := os.Getpid()
		fmt.Printf("Current PID: %d\n", pid)
		time.Sleep(2 * time.Second)
	}

}

func checkAndCreateSecret() error {
	// Check if file exists
	_, err := os.Stat("runtime.secret")
	if err == nil {
		// File exists, do nothing
		return nil
	}
	if !os.IsNotExist(err) {
		// Some other error occurred
		return err
	}

	// File doesn't exist, create it with content
	err = os.WriteFile("runtime.secret", []byte("I wrote this during running! t0P s3cret1!11 "), 0644)
	if err != nil {
		return err
	}
	return nil
}
