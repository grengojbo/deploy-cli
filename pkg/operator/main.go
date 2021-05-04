package operator

import (
	"fmt"

	"github.com/grengojbo/deploy-cli/pkg/types"
	log "github.com/sirupsen/logrus"
)

// / RunSshCommand выполнение комманд на удаленном хосте по ssh TODO: tranclate
func RunSshCommand(myCommand string, spec *types.DeploySpec, dryRun bool) error {
	ssh := SSHOperator{}
	ssh.NewSSHOperator(spec.Node)
	prepareCommand := fmt.Sprintf("%s%s%s", spec.Node.GetEnvExport(), spec.Node.GetWorkdir(), myCommand)
	log.Infof("Connection: ssh %s@%s ...", spec.Node.User, spec.Node.Host)
	if dryRun {
		log.Infof("Executing: %s\n", prepareCommand)
	} else {
		log.Debugf("Executing: %s\n", prepareCommand)
		// Выполняем комманду по SSH
		if _, err := ssh.Run(prepareCommand); err != nil {
			return err
		}
		// Demo stream
		// ssh.Stream("for i in {1..5}; do echo ${i}; sleep ; done; exit 2;", true)
		// ssh.Stream("apt update -y", false)
	}

	return nil
}
