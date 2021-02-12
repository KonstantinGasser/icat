package resources

import (
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

type NetConnSFTP struct {
	user, host, port, src string
}

func (conn *NetConnSFTP) resolve(src string) error {
	re := regexp.MustCompile("^\\S+@\\S+:[0-9]{1,6}")
	connBytes := re.Find([]byte(src))
	if connBytes == nil {
		return fmt.Errorf("could not identify the components (user,host,port) in %s", src)
	}

	userSplitSet := strings.Split(string(connBytes), "@")
	hostPortSpiltSet := strings.Split(userSplitSet[1], ":")
	imageSrc := strings.Split(src, string(connBytes))[1]

	conn.user = userSplitSet[0]
	conn.host = hostPortSpiltSet[0]
	conn.port = hostPortSpiltSet[1]
	conn.src = imageSrc
	return nil
}

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

	cleanup := func() {
		if sshConn != nil {
			sshConn.Close()
		}
		if sftpConn != nil {
			sftpConn.Close()
		}
	}
	remoteFile, err := sftpConn.Open(conn.src)
	if err != nil {
		return nil, cleanup, nil
	}
	return remoteFile, cleanup, nil
}
