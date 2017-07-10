package commands

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"

	"github.com/lcanal/opsmc/vm"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

//RunUpdates runs the package manager updates on all packages.
func RunUpdates(vms []vm.VM) {
	for _, vm := range vms {
		fmt.Printf("Connecting to VMID: %s Host: %s\n", vm.ID, vm.DNSName)
		sshConnection, err := sshConnect(vm.IP)
		if err != nil {
			fmt.Println(err)
			return
		}

		session, err := sshConnection.NewSession()
		if err != nil {
			fmt.Printf("Failed to create session: %s", err.Error())
			return
		}

		modes := ssh.TerminalModes{
			ssh.ECHO:          0,     //Disable echoing
			ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
			ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
		}
		if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
			session.Close()
			fmt.Printf("Request for pseudo terminal failed: %s", err.Error())
		}

		updateError := session.Run("sudo yum -y update")
		if updateError != nil {
			fmt.Printf("Error running yum updates: %s", updateError.Error())
			return
		}
	}
}

func sshConnect(ipaddr string) (*ssh.Client, error) {
	//sshAuthType, err := publicKeyFile("/home/lcanal/.ssh/id_rsa")
	sshAuthType, err := sshAgent()
	if err != nil {
		return nil, fmt.Errorf("sshAuthType error: %v", err)
	}

	config := &ssh.ClientConfig{
		User:            "ec2-user",
		Auth:            []ssh.AuthMethod{sshAuthType},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	connectString := fmt.Sprintf("%s:22", ipaddr)
	connection, err := ssh.Dial("tcp", connectString, config)

	if err != nil {
		return nil, fmt.Errorf("Failed to dial: %s", err)
	}

	return connection, nil
}

func publicKeyFile(file string) (ssh.AuthMethod, error) {
	buffer, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	key, err := ssh.ParsePrivateKey(buffer)
	if err != nil {
		return nil, err
	}
	return ssh.PublicKeys(key), nil
}

func sshAgent() (ssh.AuthMethod, error) {
	sshAgent, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		return nil, err
	}
	return ssh.PublicKeysCallback(agent.NewClient(sshAgent).Signers), nil
}
