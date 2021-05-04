package cmd

import (
	"github.com/grengojbo/deploy-cli/pkg/operator"
	"github.com/grengojbo/deploy-cli/pkg/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var command string

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

			host := viper.GetString("host")
			if len(host) == 0 {
				log.Fatalln("IS NOT set host")
			}
			node := &types.NodeOpts{
				User:             viper.GetString("user"),
				Host:             host,
				Password:         viper.GetString("password"),
				Workdir:          viper.GetString("workdir"),
				SSHAuthorizedKey: viper.GetString("ssh-key"),
				SshPort:          viper.GetInt32("ssh-port"),
				SSHKey:           viper.GetString("secret-ssh-key"),
				SSHPassphrase:    viper.GetString("ssh-passphrase"),
			}
			node.SetEnv(AppFlags.Environments)
			// log.Infof("ssh %s@%s", node.User, node.Host)
			// log.Infof("command: %s%s", node.GetWorkdir(), node.GetEnvExport())
			if len(command) > 0 {
				if err := operator.RunSshCommand(command, node, viper.GetBool("dry-run")); err != nil {
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
