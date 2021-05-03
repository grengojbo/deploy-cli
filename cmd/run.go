package cmd

import (
	"github.com/grengojbo/deploy-cli/pkg/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// NewCmdNode returns a new cobra command
func NewCmdRun() *cobra.Command {

	// create new cobra command
	c := &cobra.Command{
		Use:   "run",
		Short: "Run command in node(s)",
		Long:  `Run command in  node(s)`,
		Run: func(c *cobra.Command, args []string) {
			// if err := c.Help(); err != nil {
			// 	log.Errorln("Couldn't get help text")
			// 	log.Fatalln(err)
			// }
			log.Infoln("Start Run command...")
			log.Infoln("Password: ", viper.GetString("password"))
			log.Infof("ENV: %v", types.GetEnvExport(AppFlags.Environments))
		},
	}

	// add subcommands
	// cmd.AddCommand(NewCmdNodeCreate())
	// cmd.AddCommand(NewCmdNodeStart())
	// cmd.AddCommand(NewCmdNodeStop())
	// cmd.AddCommand(NewCmdNodeDelete())
	// cmd.AddCommand(NewCmdNodeList())

	// add flags

	// done
	return c
}
