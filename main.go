package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/lcanal/opsmc/commands"
)

func main() {
	var isQuiet bool
	flag.BoolVar(&isQuiet, "q", false, "Only print VIM ids, suitable to pipe into another program or this program.")
	flag.Parse()
	command := flag.Arg(0)
	if flag.NArg() == 0 {
		printHelp()
		os.Exit(0)
	}

	session := initAWSSession("us-east-1")
	ec2service := ec2.New(session)

	switch command {
	case "ls":
		vms := commands.ListVMs(ec2service)
		if !isQuiet {
			fmt.Printf("Have %d VMs running\n", len(vms))
		}
		for _, vm := range vms {
			if isQuiet {
				fmt.Println(vm.ID)
			} else {
				fmt.Printf("ID: %-30s %-30s %-20s\n", vm.ID, vm.IP, vm.DNSName)
			}
		}
	case "updates":
		commands.RunUpdates(commands.ListVMs(ec2service))
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
	fmt.Printf("%-20s%5s", " updates", "Update all packages on specific hosts.\n")
}

func initAWSSession(region string) *session.Session {
	if len(region) == 0 {
		region = "us-east-1"
	}

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if err != nil {
		log.Fatalf("Error initializing session: %s", err.Error())
	}

	return sess
}

func appUsage() {
	fmt.Printf("")
}
