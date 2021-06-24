package driver

import (
	"regexp"
)

// NewCiscoSMBDevice takes a NetDevice and initializes
// a CiscoSMBDevice.
func NewCiscoSMBDevice(d NetDevice) NetDevice {
	// Prompts
	d.Prompt.User = regexp.MustCompile(`(?im)[a-z0-9.\\-_@()/:]{1,63}>$`)
	d.Prompt.SuperUser = regexp.MustCompile(`(?im)[a-z0-9.\\-_@()/:]{1,63}#$`)
	d.Prompt.Config = regexp.MustCompile(`(?im)[a-z0-9.\-_@/:]{1,63}\([a-z0-9.\-@/:\+]{0,32}\)#$`)

	// SSH Params
	InitSSHParams(&d.SSHParams)

	// Timeout
	d.Timeout = 120

	return d
}

func CiscoSMBConnectWithSSH(d *NetDevice) error {

	clientConfig, err := SSHClientConfig(d.Credentials, d.SSHParams)
	if err != nil {
		return err
	}

	sshConn, err := ConnectWithSSH(d.IP, d.SSHParams.Port, clientConfig)
	if err != nil {
		return err
	}

	ReadSSH(sshConn.StdOut, d.Prompt.SuperUser, 5)

	d.SSHConn = sshConn

	d.SendCommandWithSSH("terminal datadump")
	d.SendCommandWithSSH("terminal width 512")

	return nil
}
