package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/lcanal/opsmc/commands"
)

func main() {
	appArgs := os.Args[1:]
	if len(appArgs) == 0 {
		printHelp()
		os.Exit(0)
	}

	session := initAWSSession("us-east-1")
	command := appArgs[0]

	switch command {
	case "ls":
		ec2service := ec2.New(session)
		vms := commands.ListVMs(ec2service)
		for _, vm := range vms {
			fmt.Printf("VMs obtained: %s w ip %s\n", vm.Name, vm.IP)
		}
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

func initAWSSession(region string) *session.Session {
	if len(region) == 0 {
		region = "us-east-1"
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		log.Fatalf("Error initializing session: %e", err)
	}

	return sess
}
