package commands

import (
	"fmt"

	"github.com/lcanal/opsmc/vm"
)

//RunUpdates runs the package manager updates on all packages.
func RunUpdates(vms []vm.VM) {
	for _, vm := range vms {
		fmt.Printf("VMID: %s\n", vm.ID)
	}
}
