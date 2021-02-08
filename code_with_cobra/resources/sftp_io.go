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
	var remoteAddr string

	user, host, imgPath, err := parseSrc(src)
	if err != nil {
		return nil, err
	}
	remoteAddr = fmt.Sprintf("%s:%d", host, c.port)

	// check if requests host is in known_hosts
	if err := c.verifyHost(host); err != nil {
		return nil, err
	}

	password, err := getPassword()
	if err != nil {
		return nil, err
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		HostKeyCallback: ssh.FixedHostKey(c.publicKey),
	}
	// set sshConn at client so ssh connection can later be closed after use
	c.sshConn, err = ssh.Dial("tcp", remoteAddr, config)
	if err != nil {
		return nil, fmt.Errorf("cloud not connect to: %s", remoteAddr)
	}

	// set sftpClient at client so client can later be closed after use
	c.sftpClient, err = sftp.NewClient(c.sshConn)
	if err != nil {
		return nil, fmt.Errorf("cloud not create SFTP-Client")
	}

	remoteFile, err := c.sftpClient.Open(imgPath)
	if err != nil {
		return nil, fmt.Errorf("cloud not open remote file: %s", imgPath)
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

func parseSrc(src string) (user, host, imgPath string, err error) {
	// src is tagen from os.Args[2] -> if --remote set before the src
	// os.Args[2] == --remote
	if strings.Contains(src, "--remote") {
		err = fmt.Errorf("--remote flag must be set at the end of the command (icat view <src> --remote)")
		return "", "", "", err
	}

	// user@server:/path/to/image -> user@server, /path/to/image
	split := strings.Split(src, ":")

	// parese user, server: user@server -> user, server
	credentials := strings.Split(split[0], "@")
	if len(credentials) < 2 {
		err = fmt.Errorf("you need to provide both, user and server in the form of user@server")
		return "", "", "", err
	}
	user, host = credentials[0], credentials[1]

	// path of image can containe colon i.ex /path/to:some/image
	// join path back together (split created due to first spilit)
	imgPath = split[1]
	if len(split[1:]) > 1 {
		imgPath = strings.Join(split[1:], ":")
	}

	return user, host, imgPath, nil
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
