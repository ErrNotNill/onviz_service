package fstp

import (
	"bufio"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"net"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

var (
	RootSftpUrl = os.Getenv("SFTP_ROOT_URL")
	SFTP        = os.Getenv("SFTP_URL")
	HostUrl     = os.Getenv("SFTP_HOST_URL")
)

func StartSfTp() {
	// Get SFTP To Go URL from environment
	//rawurl := os.Getenv("SFTPTOGO_URL")

	fmt.Printf("%q", SFTP)
	fmt.Println()
	urlStr := `testsftp ` + RootSftpUrl
	parsedUrl, err := url.Parse(urlStr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse SFTP To Go URL: %s\n", err)
		os.Exit(1)
	}

	user := parsedUrl.User.Username()
	pass, _ := parsedUrl.User.Password()

	// Default SFTP port
	port := 22

	hostKey := getHostKey(HostUrl)

	fmt.Fprintf(os.Stdout, "Connecting to %s ...\n", HostUrl)

	var auths []ssh.AuthMethod

	// Try to use $SSH_AUTH_SOCK which contains the path of the unix file socket that the sshd agent uses
	// for communication with other processes.
	if aconn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		auths = append(auths, ssh.PublicKeysCallback(agent.NewClient(aconn).Signers))
	}

	// Use password authentication if provided
	if pass != "" {
		auths = append(auths, ssh.Password(pass))
	}

	// Initialize client configuration
	config := ssh.ClientConfig{
		User: user,
		Auth: auths,
		// Uncomment to ignore host key check
		//HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		HostKeyCallback: ssh.FixedHostKey(hostKey),
	}

	addr := fmt.Sprintf("%s:%d", HostUrl, port)

	// Connect to server
	conn, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to connecto to [%s]: %v\n", addr, err)
		os.Exit(1)
	}

	defer conn.Close()

	// Create new SFTP client
	sc, err := sftp.NewClient(conn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to start SFTP subsystem: %v\n", err)
		os.Exit(1)
	}
	defer sc.Close()
}

// Get host key from local known hosts
func getHostKey(host string) ssh.PublicKey {
	// parse OpenSSH known_hosts file
	// ssh or use ssh-keyscan to get initial key
	file, err := os.Open(filepath.Join(os.Getenv("HOME"), ".ssh", "known_hosts"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to read known_hosts file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var hostKey ssh.PublicKey
	for scanner.Scan() {
		fields := strings.Split(scanner.Text(), " ")
		if len(fields) != 3 {
			continue
		}
		if strings.Contains(fields[0], host) {
			var err error
			hostKey, _, _, _, err = ssh.ParseAuthorizedKey(scanner.Bytes())
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing %q: %v\n", fields[2], err)
				os.Exit(1)
			}
			break
		}
	}

	if hostKey == nil {
		fmt.Fprintf(os.Stderr, "No hostkey found for %s", host)
		os.Exit(1)
	}

	return hostKey
}

func listFiles(sc sftp.Client, remoteDir string) (err error) {
	fmt.Fprintf(os.Stdout, "Listing [%s] ...\n\n", remoteDir)

	files, err := sc.ReadDir(remoteDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to list remote dir: %v\n", err)
		return
	}

	for _, f := range files {
		var name, modTime, size string

		name = f.Name()
		modTime = f.ModTime().Format("2006-01-02 15:04:05")
		size = fmt.Sprintf("%12d", f.Size())

		if f.IsDir() {
			name = name + "/"
			modTime = ""
			size = "PRE"
		}
		// Output each file name and size in bytes
		fmt.Fprintf(os.Stdout, "%19s %12s %s\n", modTime, size, name)
	}

	return
}
