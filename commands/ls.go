package commands

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/lcanal/opsmc/vm"
)

//ListVMs List a set of VMs, taking in a cloud initiated service.
func ListVMs(service *ec2.EC2) []vm.VM {
	vmList := make([]vm.VM, 0)

	params := &ec2.DescribeInstancesInput{
		Filters: []*ec2.Filter{
			{
				Name:   aws.String("instance-state-name"),
				Values: []*string{aws.String("running"), aws.String("pending")},
			},
		},
	}

	resp, err := service.DescribeInstances(params)

	if err != nil {
		log.Fatalf("There was an error listing instances in %e", err)
	}

	for idx := range resp.Reservations {
		for _, inst := range resp.Reservations[idx].Instances {
			var foundVM vm.VM
			if inst.PublicIpAddress == nil {
				foundVM.IP = "No IP Address Assigned"
			} else {
				foundVM.IP = *inst.PublicIpAddress
			}
			if inst.PublicDnsName == nil {
				foundVM.DNSName = "No DNS Name Assigned"
			} else {
				foundVM.DNSName = *inst.PublicDnsName
			}

			foundVM.ID = *inst.InstanceId
			foundVM.Type = *inst.InstanceType
			foundVM.FullStatus = inst.GoString()

			vmList = append(vmList, foundVM)
		}
	}

	return vmList
}
