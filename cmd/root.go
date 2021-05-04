/*
Copyright © 2020 The k3d Author(s)

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
package cmd

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cliutil "github.com/grengojbo/deploy-cli/cmd/util"
	"github.com/grengojbo/deploy-cli/pkg/types"
	"github.com/grengojbo/deploy-cli/version"
	log "github.com/sirupsen/logrus"
	"github.com/sirupsen/logrus/hooks/writer"
)

// RootFlags describes a struct that holds flags that can be set on root level of the command
type RootFlags struct {
	debugLogging       bool
	traceLogging       bool
	timestampedLogging bool
	version            bool
	cfgFile            string
	Environments       []string
}

var AppFlags = RootFlags{}

// var envs = []string{}
var Verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "deploy-cli",
	Short: "Deploy CLI",
	Long: `
Deploy wrapper CLI that helps you to easily git, dockec, podman, docker-compose, make.
`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("Start...")
		if AppFlags.version {
			printVersion()
		} else {
			if err := cmd.Usage(); err != nil {
				log.Fatalln(err)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if len(os.Args) > 1 {
		parts := os.Args[1:]
		// Check if it's a built-in command, else try to execute it as a plugin
		if _, _, err := rootCmd.Find(parts); err != nil {
			pluginFound, err := cliutil.HandlePlugin(context.Background(), parts)
			if err != nil {
				log.Errorf("Failed to execute plugin '%+v'", parts)
				log.Fatalln(err)
			} else if pluginFound {
				os.Exit(0)
			}
		}
	}
	if err := rootCmd.Execute(); err != nil {
		log.Fatalln(err)
	}
}

func init() {

	// defaultConfigName = "env"
	// cobra.OnInitialize(initConfig)
	// Init
	cobra.OnInitialize(initLogging, initConfig)
	// cobra.OnInitialize(initLogging, initRuntime)

	rootCmd.PersistentFlags().BoolVar(&AppFlags.debugLogging, "verbose", false, "Enable verbose output (debug logging)")
	rootCmd.PersistentFlags().BoolVar(&AppFlags.traceLogging, "trace", false, "Enable super verbose output (trace logging)")
	rootCmd.PersistentFlags().BoolVar(&AppFlags.timestampedLogging, "timestamps", false, "Enable Log timestamps")

	rootCmd.PersistentFlags().String("host", "", "Public hostname of node on which to run command")
	_ = viper.BindPFlag("host", rootCmd.PersistentFlags().Lookup("host"))

	rootCmd.PersistentFlags().StringP("password", "p", "", "SSH password")
	_ = viper.BindPFlag("password", rootCmd.PersistentFlags().Lookup("password"))
	_ = viper.BindEnv("password", "SECRET_SSH_PASSWORD")

	rootCmd.PersistentFlags().StringP("user", "u", "root", "Username for SSH login")
	_ = viper.BindPFlag("user", rootCmd.PersistentFlags().Lookup("user"))
	_ = viper.BindEnv("user", "SECRET_SSH_USERNAME")

	rootCmd.PersistentFlags().String("ssh-key", "~/.ssh/id_rsa", "The ssh key to use for remote login")
	_ = viper.BindPFlag("ssh-key", rootCmd.PersistentFlags().Lookup("ssh-key"))

	rootCmd.PersistentFlags().Int("ssh-port", 22, "The port on which to connect for ssh")
	_ = viper.BindPFlag("ssh-port", rootCmd.PersistentFlags().Lookup("ssh-port"))

	// rootCmd.PersistentFlags().Bool("sudo", true, "Use sudo for installation. e.g. set to false when using the root user and no sudo is available.")
	// _ = viper.BindPFlag("sudo", rootCmd.PersistentFlags().Lookup("sudo"))

	rootCmd.PersistentFlags().StringSliceVar(&AppFlags.Environments, "set-env", []string{}, "Set environment variable")

	rootCmd.PersistentFlags().StringP("env", "e", "", "Set environment developer, testing, staging, production")
	_ = viper.BindPFlag("env", rootCmd.PersistentFlags().Lookup("env"))

	// cmd.Flags().StringArrayP("port", "p", nil, "Map ports from the node containers to the host (Format: `[HOST:][HOSTPORT:]CONTAINERPORT[/PROTOCOL][@NODEFILTER]`)\n - Example: `k3d cluster create --agents 2 -p 8080:80@agent[0] -p 8081@agent[1]`")
	// _ = ppViper.BindPFlag("cli.ports", cmd.Flags().Lookup("port"))

	// viper.SetDefault("port", 80)

	rootCmd.PersistentFlags().StringP("workdir", "w", "", "Set workdir")
	_ = viper.BindPFlag("workdir", rootCmd.PersistentFlags().Lookup("workdir"))

	rootCmd.PersistentFlags().Bool("dry-run", false, "Show run command")
	_ = viper.BindPFlag("dry-run", rootCmd.PersistentFlags().Lookup("dry-run"))
	viper.SetDefault("secret-ssh-key", "")
	_ = viper.BindEnv("secret-ssh-key", "SECRET_SSH_KEY")
	_ = viper.BindEnv("ssh-passphrase", "SECRET_SSH_PASSPHRASE")
	viper.SetDefault("disable-run-command", false)
	_ = viper.BindEnv("disable-run-command", "SECRET_DEPLOY_DISABLE_RUN")

	// add local flags
	rootCmd.Flags().BoolVar(&AppFlags.version, "version", false, "Show deploy-cli version")

	// add subcommands
	rootCmd.AddCommand(NewCmdCompletion())
	rootCmd.AddCommand(NewCmdRun())
	// rootCmd.AddCommand(kubeconfig.NewCmdKubeconfig())
	// rootCmd.AddCommand(node.NewCmdNode())
	// rootCmd.AddCommand(image.NewCmdImage())
	// rootCmd.AddCommand(cfg.NewCmdConfig())
	// rootCmd.AddCommand(registry.NewCmdRegistry())

	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "Show deploy-cli version",
		Long:  "Show deploy-cli version",
		Run: func(cmd *cobra.Command, args []string) {
			printVersion()
		},
	})

}

// initConfig читает в файле конфигурации и переменных ENV, если они установлены.
func initConfig() {

	if AppFlags.cfgFile != "" {
		// Use config file from the flag.
		viper.AddConfigPath("$HOME")
		viper.AddConfigPath(".")
		viper.SetConfigFile(AppFlags.cfgFile)

	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".bboxApi" (without extension).
		// viper.AddConfigPath(home)
		// viper.AddConfigPath(".")
		// viper.SetConfigName(defaultConfigName)
		log.Debugln("Home: ", home)
		// fmt.Println("------------- initConfig -----------", home)
	}

	viper.SetEnvPrefix(types.DefaultEnvPrefix)
	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		if Verbose {
			fmt.Println("Using config file:", viper.ConfigFileUsed())
		}

		// viper.WatchConfig()
		// viper.OnConfigChange(func(e fsnotify.Event) {
		// 	if Verbose {
		// 		fmt.Println("Config file changed:", e.Name)
		// 	}
		// })
	}
}

// initLogging initializes the logger
func initLogging() {
	if AppFlags.traceLogging {
		log.SetLevel(log.TraceLevel)
	} else if AppFlags.debugLogging {
		log.SetLevel(log.DebugLevel)
	} else {
		switch logLevel := strings.ToUpper(os.Getenv("LOG_LEVEL")); logLevel {
		case "TRACE":
			log.SetLevel(log.TraceLevel)
		case "DEBUG":
			log.SetLevel(log.DebugLevel)
		case "WARN":
			log.SetLevel(log.WarnLevel)
		case "ERROR":
			log.SetLevel(log.ErrorLevel)
		default:
			log.SetLevel(log.InfoLevel)
		}
	}
	log.SetOutput(ioutil.Discard)
	log.AddHook(&writer.Hook{
		Writer: os.Stderr,
		LogLevels: []log.Level{
			log.PanicLevel,
			log.FatalLevel,
			log.ErrorLevel,
			log.WarnLevel,
		},
	})
	log.AddHook(&writer.Hook{
		Writer: os.Stdout,
		LogLevels: []log.Level{
			log.InfoLevel,
			log.DebugLevel,
			log.TraceLevel,
		},
	})

	formatter := &log.TextFormatter{
		ForceColors: true,
	}

	if AppFlags.timestampedLogging || os.Getenv("LOG_TIMESTAMPS") != "" {
		formatter.FullTimestamp = true
	}

	log.SetFormatter(formatter)

}

func printVersion() {
	fmt.Printf("deploy-cli version %s\n", version.GetVersion())
}

func generateFishCompletion(writer io.Writer) error {
	return rootCmd.GenFishCompletion(writer, true)
}

// Completion
var completionFunctions = map[string]func(io.Writer) error{
	"bash": rootCmd.GenBashCompletion,
	"zsh": func(writer io.Writer) error {
		if err := rootCmd.GenZshCompletion(writer); err != nil {
			return err
		}

		fmt.Fprintf(writer, "\n# source completion file\ncompdef _deploy-cli deploy-cli\n")

		return nil
	},
	"psh":        rootCmd.GenPowerShellCompletion,
	"powershell": rootCmd.GenPowerShellCompletion,
	"fish":       generateFishCompletion,
}

// NewCmdCompletion creates a new completion command
func NewCmdCompletion() *cobra.Command {
	// create new cobra command
	cmd := &cobra.Command{
		Use:   "completion SHELL",
		Short: "Generate completion scripts for [bash, zsh, fish, powershell | psh]",
		Long:  `Generate completion scripts for [bash, zsh, fish, powershell | psh]`,
		Args:  cobra.ExactArgs(1), // TODO: NewCmdCompletion: add support for 0 args = auto detection
		Run: func(cmd *cobra.Command, args []string) {
			if completionFunc, ok := completionFunctions[args[0]]; ok {
				if err := completionFunc(os.Stdout); err != nil {
					log.Fatalf("Failed to generate completion script for shell '%s'", args[0])
				}
				return
			}
			log.Fatalf("Shell '%s' not supported for completion", args[0])
		},
	}
	return cmd
}
