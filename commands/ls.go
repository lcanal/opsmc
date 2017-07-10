package commands

import (
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/lcanal/opsmc/vm"
)

//ListVMs List a set of VMs, taking in a cloud initiated service.
func ListVMs(service *ec2.EC2) []vm.VM {
	vmList := make([]vm.VM, 0)
	return vmList
}
