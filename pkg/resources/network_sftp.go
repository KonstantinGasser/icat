package resources

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

const (
	defaultPort = "22"
	// regexUserHostPort regex for: user@host:port
	regexUserHostPort = "(^\\w+@\\w+:[0-9]{1,1000})"
	// regexUserHost regex for: user@host
	regexUserHost = "(^\\w+@[a-zA-Z]*)"
	//	regexUserIP regex for: user@ip
	regexUserIP = "(^\\w+@(?:[0-9]{1,3}\\.){3}[0-9]{1,3})"
	// regexUserIPPort regex for: user@ip:port
	regexUserIPPort = "(^\\w+@(?:[0-9]{1,3}\\.){3}[0-9]{1,3}:[0-9]{1,1000})"
)

// NetConnSFTP handles images requested from a remote sever
type NetConnSFTP struct {
	user, host, port, src string
}

func (conn *NetConnSFTP) resolve(src string) error {
	// match user@hostname:port | user@ip.address:port in src string
	re := regexp.MustCompile(fmt.Sprintf("%s|%s|%s|%s",
		regexUserHostPort,
		regexUserIPPort,
		regexUserIP,
		regexUserHost,
	))
	connBytes := re.Find([]byte(src))
	if connBytes == nil {
		return fmt.Errorf("could not identify the components (user,host,port) in %s", src)
	}
	// -> [user, host:port]
	userSplitSet := strings.Split(string(connBytes), "@")
	// -> [host, port]
	hostPortSpiltSet := strings.Split(userSplitSet[1], ":")
	// -> [user@host:port, /path/to/image]
	imageSrc := strings.Split(src, string(connBytes))[1]

	conn.user = userSplitSet[0]
	conn.host = hostPortSpiltSet[0]
	conn.port = defaultPort
	if len(hostPortSpiltSet) > 1 {
		conn.port = hostPortSpiltSet[1]
	}
	conn.src = imageSrc
	return nil
}

// Open opens a sftp connection to read a file from a remote server
// returns a cleanup function to close the ssh.Connection and sftp.Connection
func (conn NetConnSFTP) Open() (io.Reader, func(), error) {

	sshConfig, err := getSFTPConfig(conn.user, conn.host)
	if err != nil {
		return nil, nil, err
	}
	sshConn, err := ssh.Dial("tcp", fmt.Sprintf("%s:%s", conn.host, conn.port), sshConfig)
	if err != nil {
		return nil, nil, err
	}

	sftpConn, err := sftp.NewClient(sshConn)
	if err != nil {
		return nil, func() {
			if sshConn != nil {
				sshConn.Close()
			}
		}, err
	}

	// close sshConn, sftpConn if opening file fails
	cleanupConns := func() {
		if sftpConn != nil {
			sftpConn.Close()
		}
		if sshConn != nil {
			sshConn.Close()
		}
	}
	remoteFile, err := sftpConn.Open(conn.src)
	if err != nil {
		return nil, cleanupConns, nil
	}

	// close sshConn, sftpConn and opened file
	// after request has been served
	cleanupAll := func() {
		if remoteFile != nil {
			remoteFile.Close()
		}
		if sftpConn != nil {
			sftpConn.Close()
		}
		if sshConn != nil {
			sshConn.Close()
		}
	}
	return remoteFile, cleanupAll, nil
}
