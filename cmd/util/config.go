package util

import (
	"strings"

	"github.com/grengojbo/deploy-cli/pkg/config"
	"github.com/grengojbo/deploy-cli/pkg/types"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var cfgViper = viper.New()

// FromViper load config from viper
func FromViper(args []string) (spec *types.DeploySpec, err error) {
	spec = &types.DeploySpec{
		Node:    &types.NodeOpts{},
		Bastion: &types.NodeOpts{},
	}
	if len(args) > 0 {

		cfgViper = config.GetViperConfig(args[0], strings.ToLower(viper.GetString("env")))
		// var cfg types.NodeOpts

		// determine config kind
		if cfgViper.GetString("kind") != "" && strings.ToLower(cfgViper.GetString("kind")) != "simple" {
			log.Fatalf("Wrong `kind` '%s' != 'simple' in config file", cfgViper.GetString("kind"))
		}
		if err := cfgViper.Unmarshal(&spec); err != nil {
			log.Errorln("Failed to unmarshal File config")
			return spec, err
		}
	}

	if len(viper.GetString("user")) > 0 {
		log.Debugln("Replace user:", viper.GetString("user"))
		spec.Node.User = viper.GetString("user")
	} else if len(spec.Node.User) == 0 {
		spec.Node.User = "root"
	}

	if len(viper.GetString("password")) > 0 {
		log.Debugln("Replace password: ********")
		spec.Node.Password = viper.GetString("password")
	}

	if viper.GetString("ssh-key") != string(types.DefaultSSHKeyPath) {
		log.Debugln("Replace ssh-key:", viper.GetString("ssh-key"))
		spec.Node.SSHAuthorizedKey = viper.GetString("ssh-key")
	}
	if len(spec.Node.SSHAuthorizedKey) == 0 {
		spec.Node.SSHAuthorizedKey = string(types.DefaultSSHKeyPath)
		log.Debugln("Replace ssh-key: ", spec.Node.SSHAuthorizedKey)
	}

	if viper.GetInt32("ssh-port") != int32(types.DefaultSshPort) {
		log.Debugln("Replace ssh-port:", viper.GetInt32("ssh-port"))
		spec.Node.SshPort = viper.GetInt32("ssh-port")
	}
	if spec.Node.SshPort == 0 {
		spec.Node.SshPort = int32(types.DefaultSshPort)
		log.Debugln("Replace ssh-port: ", spec.Node.SshPort)
	}

	if len(viper.GetString("host")) > 0 {
		log.Debugln("Replace host:", viper.GetString("host"))
		spec.Node.Host = viper.GetString("host")
	}
	if len(spec.Node.Host) == 0 {
		log.Fatalln("IS NOT set host")
	}

	if len(viper.GetString("workdir")) > 0 {
		log.Debugln("Replace workdir:", viper.GetString("workdir"))
		spec.Node.Workdir = viper.GetString("workdir")
	}

	if len(viper.GetString("secret-ssh-key")) > 0 {
		log.Debugln("Replace source content of private key: ****")
		spec.Node.SSHKey = viper.GetString("secret-ssh-key")
	}
	if len(viper.GetString("ssh-passphrase")) > 0 {
		log.Debugln("Replace PrivateKey With Passphrase: ****")
		spec.Node.SSHPassphrase = viper.GetString("ssh-passphrase")
	}

	return spec, nil
}

// node = &types.NodeOpts{
// 					Host:             host,
// 					Workdir:          viper.GetString("workdir"),
// SSHKey:           viper.GetString("secret-ssh-key"),
// 					SSHPassphrase:    viper.GetString("ssh-passphrase"),
// 				}
