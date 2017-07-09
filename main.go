package main

import (
	"fmt"
	"os"
)

func main() {
	appArgs := os.Args[1:]
	if len(appArgs) == 0 {
		printHelp()
		os.Exit(0)
	}

	command := appArgs[0]
	//commandOps := appArgs[1:]
	switch command {
	case "ls":
		fmt.Println("vms ls ls ls vms ls lslsls")

	default:
		fmt.Printf("Error: Unknown command %s\n", command)
	}
}

func printHelp() {
	fmt.Printf("Usage: %s command [options]\n", os.Args[0])
	fmt.Printf("Available Commands\n")
	fmt.Printf("%-20s%5s", " ls", "List your currently running EC2 instances\n")
	fmt.Printf("%-20s%5s", " mk", "Make a new EC2 instance\n")
	fmt.Printf("%-20s%5s", " rm", "Remove a specific AWS EC2 instance.\n")
}
