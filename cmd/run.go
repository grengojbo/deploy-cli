package cmd

import (
	"strings"

	"github.com/grengojbo/deploy-cli/cmd/util"
	"github.com/grengojbo/deploy-cli/pkg/operator"
	"github.com/grengojbo/deploy-cli/pkg/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var command string
var spec *types.DeploySpec

// var node *types.NodeOpts
var err error

// NewCmdNode returns a new cobra command
func NewCmdRun() *cobra.Command {

	// create new cobra command
	c := &cobra.Command{
		Use:   "run",
		Short: "Run command in node(s)",
		Long:  `Run command in  node(s)`,
		Args:  cobra.RangeArgs(0, 1),
		Run: func(c *cobra.Command, args []string) {
			// if err := c.Help(); err != nil {
			// 	log.Errorln("Couldn't get help text")
			// 	log.Fatalln(err)
			// }
			if viper.GetBool("disable-run-command") {
				log.Fatalln("RUN command disable. Please read the documentation https://zerobox.atlassian.net/wiki/external/10420225/MzhiNWJmMzUyY2Y0NGI0ZThlNTdhYmNlYmI1ZTJjMjU?atlOrigin=eyJpIjoiOTlhODQ1NTkwMDJkNDVhZGE5MTc0Y2UyMmZkZTUwNTUiLCJwIjoiYyJ9")
			}
			log.Infoln("Start Run command...")
			if len(viper.GetString("env")) > 0 {
				log.Infof("environments: %s", strings.ToLower(viper.GetString("env")))
			}

			spec, err = util.FromViper(args)
			if err != nil {
				log.Errorln(err.Error())
			}
			// if len(args) == 0 {
			// 	if len(viper.GetString("host")) == 0 {
			// 		log.Fatalln("IS NOT set host")
			// 	}
			// 	spec.Node.Host = viper.GetString("host")
			// }

			spec.Node.SetEnv(AppFlags.Environments)
			// log.Infof("ssh %s@%s", node.User, node.Host)
			// log.Infof("command: %s%s", node.GetWorkdir(), node.GetEnvExport())
			if len(command) > 0 {
				if err := operator.RunSshCommand(command, spec, viper.GetBool("dry-run")); err != nil {
					log.Fatal(err.Error())
				}
			} else {
				log.Errorln("is not run command :(")
			}
		},
	}

	// add subcommands
	// cmd.AddCommand(NewCmdNodeCreate())
	// cmd.AddCommand(NewCmdNodeStart())
	// cmd.AddCommand(NewCmdNodeStop())
	// cmd.AddCommand(NewCmdNodeDelete())
	// cmd.AddCommand(NewCmdNodeList())

	// add flags
	c.Flags().StringVarP(&command, "command", "c", "", "exec command")
	// cmd.Flags().StringP("image", "i", "", "Specify k3s image that you want to use for the nodes")
	// _ = cfgViper.BindPFlag("image", cmd.Flags().Lookup("image"))
	// cfgViper.SetDefault("image", fmt.Sprintf("%s:%s", k3d.DefaultK3sImageRepo, version.GetK3sVersion(false)))
	// done
	return c
}
