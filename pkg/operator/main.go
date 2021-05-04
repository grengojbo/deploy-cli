package operator

import (
	"fmt"

	"github.com/grengojbo/deploy-cli/pkg/types"
	log "github.com/sirupsen/logrus"
)

// / RunSshCommand выполнение комманд на удаленном хосте по ssh TODO: tranclate
func RunSshCommand(myCommand string, node *types.NodeOpts, dryRun bool) error {
	ssh := SSHOperator{}
	ssh.NewSSHOperator(node)
	prepareCommand := fmt.Sprintf("%s%s%s", node.GetEnvExport(), node.GetWorkdir(), myCommand)
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
