package resources

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

func getSFTPConfig(user, host string) (*ssh.ClientConfig, error) {

	pubKey, err := getPubKey(host)
	if err != nil {
		return nil, err
	}

	// parse server password from the os.Stdin
	fmt.Print("Server Password: ")
	defer fmt.Println()
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return nil, fmt.Errorf("could not read password from os.Stdin")
	}
	password := strings.TrimSpace(string(bytePassword))

	return &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.FixedHostKey(pubKey),
	}, nil
}

func getPubKey(host string) (pubKey ssh.PublicKey, err error) {

	knownFile, err := os.Open(path.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		return nil, err
	}
	defer knownFile.Close()

	scanner := bufio.NewScanner(knownFile)

	for scanner.Scan() {
		tags := strings.Split(scanner.Text(), " ")
		if len(tags) != 3 {
			continue
		}
		if strings.Contains(tags[0], host) {
			pubKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				return nil, fmt.Errorf("could not parse ssh public key for: %s :%s", tags[2], err.Error())
			}
		}
	}
	return pubKey, nil
}
