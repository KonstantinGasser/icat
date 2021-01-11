package resources

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"syscall"

	"github.com/KonstantinGasser/sherlocked/cmd_errors"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

type sftpclient struct {
	port       int
	sshConn    *ssh.Client
	sftpClient *sftp.Client
	publicKey  ssh.PublicKey
}

func NewSFTP(port int) Resource {
	return &sftpclient{
		port: port,
	}
}

func (c *sftpclient) Open(src string) (io.Reader, error) {
	var imgPath, user, host, remoteAddr string

	// TODO: peforme error checking for src string
	// spilt user@server:/path/to/image -> user@server, /path/to/image
	parts := strings.Split(src, ":")
	
	// split user@server -> user, server
	credentials := strings.Split(parts[0], "@")
	user, host = credentials[0], credentials[1]
	remoteAddr = fmt.Sprintf("%s:%d", host, c.port)
	
	// TODO: put in extra func test with test cases
	// path can contain : /path/to:some/image -> will parts will then contain [..., /path/to, some/image,...]
	// join spilt fields in path back again
	imgPath = parts[1]
	if len(parts[1:]) > 1 { 
		imgPath = path.Join(parts[1:]...) // this is wrong!!!
	}

	// check if requests host is in known_hosts
	if err := c.verifyHost(credentials[1]); err != nil {
		return nil, err
	}

	pass, err := getPassword()
	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.FixedHostKey(c.publicKey),
	}
	// set sshConn at client so ssh connection can later be closed after use
	c.sshConn, err = ssh.Dial("tcp", remoteAddr, config)
	if err != nil {
		return nil, err
	}
	
	// set sftpClient at client so client can later be closed after use
	c.sftpClient, err = sftp.NewClient(c.sshConn)
	if err != nil {
		return nil, err
	}

	remoteFile, err := c.sftpClient.Open(imgPath)
	if err != nil {
		return nil, err
	}

	return remoteFile, err

}

func (c *sftpclient) Close() {
	c.sshConn.Close()
	c.sftpClient.Close()
}

func (c *sftpclient) verifyHost(host string) error {

	// open ssh known_host file
	knownHost, err := os.Open(path.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(knownHost)

	for scanner.Scan() {
		tags := strings.Split(scanner.Text(), " ")
		if len(tags) != 3 {
			continue
		}

		if strings.Contains(tags[0], host) {
			var err error
			c.publicKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				return fmt.Errorf("could not parse ssh public key for: %s :%s", tags[2], err.Error())
			}
			return nil
		}
	}

	return fmt.Errorf("could not find any known host in $HOME/.ssh/known_hosts for %s", host)
}

// TODO: move to StdOu Resource
// Password handels users password input
func getPassword() (string, error) {

	fmt.Print("Server Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", cmd_errors.OSStdInError{
			MSG: `🤨 Sorry ~ somehow we failed to read your password from os.Stdin..`,
		}
	}
	password := string(bytePassword)
	fmt.Print("\n")
	return strings.TrimSpace(password), nil
}
