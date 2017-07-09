package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		printHelp()
		os.Exit(0)
	}
}

func printHelp() {
	fmt.Printf("Usage: %s command [options]\n", os.Args[0])
	fmt.Printf("Available Commands\n")
	fmt.Printf("%-20s%5s", " ls", "List your currently running EC2 instances\n")
	fmt.Printf("%-20s%5s", " mk", "Make a new EC2 instance\n")
	fmt.Printf("%-20s%5s", " rm", "Remove a specific AWS EC2 instance.\n")
}
