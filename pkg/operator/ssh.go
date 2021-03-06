package operator

import (
	"fmt"
	"time"

	"github.com/appleboy/easyssh-proxy"
	"github.com/grengojbo/deploy-cli/pkg/types"
	"github.com/grengojbo/deploy-cli/pkg/util"
	log "github.com/sirupsen/logrus"
)

type SSHOperator struct {
	Config *easyssh.MakeConfig
}

// NewSshConnection New SSH Connection
func (r *SSHOperator) NewSSHOperator(node *types.NodeOpts) {
	r.Config = &easyssh.MakeConfig{
		User: node.User,
		// Optional key or Password without either we try to contact your agent SOCKET
		// Password: "password",
		// Paste your source content of private key
		// Key: `-----BEGIN RSA PRIVATE KEY-----
		// .........................
		// -----END RSA PRIVATE KEY-----
		// `,
		Port:    fmt.Sprintf("%d", node.SshPort),
		Timeout: 60 * time.Second,

		// Parse PrivateKey With Passphrase
		// Passphrase: "XXXX",

		// Optional fingerprint SHA256 verification
		// Get Fingerprint: ssh.FingerprintSHA256(key)
		// Fingerprint: "SHA256:................E"

		// Enable the use of insecure ciphers and key exchange methods.
		// This enables the use of the the following insecure ciphers and key exchange methods:
		// - aes128-cbc
		// - aes192-cbc
		// - aes256-cbc
		// - 3des-cbc
		// - diffie-hellman-group-exchange-sha256
		// - diffie-hellman-group-exchange-sha1
		// Those algorithms are insecure and may allow plaintext data to be recovered by an attacker.
		// UseInsecureCipher: true,
	}
	r.Config.Server = node.Host
	if len(node.Password) > 0 {
		r.Config.Password = node.Password
		log.Debugln("Set password: ****")
	} else if len(node.SSHKey) > 0 {
		r.Config.Key = node.SSHKey
		log.Debugln("Set source content of private key: ****")
	} else if len(node.SSHAuthorizedKey) > 0 {
		r.Config.KeyPath = util.ExpandPath(node.SSHAuthorizedKey)
		log.Debugf("sshKeyPath: %s", r.Config.KeyPath)
	}
	if len(node.SSHPassphrase) > 0 {
		r.Config.Passphrase = node.SSHPassphrase
		log.Debugln("Set PrivateKey With Passphrase: ****")
	}
	// log.Debugf("ssh -i %s %s@%s:%s", ssh.KeyPath, ssh.User, ssh.Server, ssh.Port)
}

// Run command on remote machine
//   Example:
func (r *SSHOperator) Run(command string) (done bool, err error) {
	stdOut, stdErr, done, err := r.Config.Run(command, 60*time.Second)
	if len(stdOut) > 0 {
		log.Debugln("===== stdOut ======")
		log.Debugf("%v", stdOut)
		log.Debugln("===================")
	}
	if len(stdErr) > 0 {
		log.Errorln("===== stdErr ======")
		log.Errorf("%v", stdErr)
		log.Errorln("===================")
	}
	return done, err
}

// Stream returns one channel that combines the stdout and stderr of the command
// as it is run on the remote machine, and another that sends true when the
// command is done. The sessions and channels will then be closed.
//  isPrint - ???????????????? ?????????????????? ???? ?????????? ?????? ?? ??????
func (r *SSHOperator) Stream(command string, isPrint bool) {
	// Call Run method with command you want to run on remote server.
	stdoutChan, stderrChan, doneChan, errChan, err := r.Config.Stream(command, 60*time.Second)
	// Handle errors
	if err != nil {
		log.Fatalln("Can't run remote command: " + err.Error())
	} else {
		// read from the output channel until the done signal is passed
		isTimeout := true
	loop:
		for {
			select {
			case isTimeout = <-doneChan:
				break loop
			case outline := <-stdoutChan:
				if isPrint && len(outline) > 0 {
					// fmt.Println("out:", outline)
					fmt.Println(outline)
				} else if len(outline) > 0 {
					log.Infoln(outline)
				}
			case errline := <-stderrChan:
				if isPrint && len(errline) > 0 {
					// fmt.Println("err:", errline)
					fmt.Println(errline)
				} else if len(errline) > 0 {
					log.Errorln(errline)
				}
			case err = <-errChan:
			}
		}

		// get exit code or command error.
		if err != nil {
			log.Errorln("Error: " + err.Error())
		}

		// command time out
		if !isTimeout {
			log.Errorln("Error: command timeout")
		}
	}
}
