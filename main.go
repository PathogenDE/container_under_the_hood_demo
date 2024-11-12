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
		// something bad happened
		return err
	}

	hostname, err := os.Hostname()
	if err != nil {
		return fmt.Errorf("failed to get hostname: %v", err)
	}

	// time and
	currentTime := time.Now().Format(time.RFC3339)
	content := fmt.Sprintf("I wrote this during running\nTime: %s\nHostname: %s\n",
		currentTime,
		hostname)

	err = os.WriteFile("runtime.secret", []byte(content), 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %v", err)
	}
	return nil
}
