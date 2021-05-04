package config

import (
	"github.com/grengojbo/deploy-cli/pkg/util"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v2"
)

var CfgViper = viper.New()

func InitConfig(cFile string, env string) {

	// cfgViper.SetEnvPrefix("K3S")
	// cfgViper.AutomaticEnv()

	// cfgViper.SetConfigType("yaml")

	configFile := util.GerConfigFileName(cFile, env)
	CfgViper.SetConfigFile(configFile)

	if err := CfgViper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Fatalf("Config file %s not found: %+v", configFile, err)
		}
		// config file found but some other error happened
		log.Fatalf("Failed to read config file %s: %+v", configFile, err)
	}

	log.Infof("Using config file %s", CfgViper.ConfigFileUsed())

	if log.GetLevel() >= log.DebugLevel {
		c, _ := yaml.Marshal(CfgViper.AllSettings())
		log.Debugf("Configuration:\n%s", c)

	}
}
