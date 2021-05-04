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
package types

import (
	"fmt"
	"strings"
	"time"
)

// DefaultRegistryImageRepo defines the default image used for the k3d-managed registry
const DefaultRegistryImageRepo = "docker.io/library/registry"

// DefaultRegistryImageTag defines the default image tag used for the k3d-managed registry
const DefaultRegistryImageTag = "2"

const GetDeployCliScript = "curl -sfL https://github.com/grengojbo/deploy-cli"

// const (
// 	MetalLb              string = "MetalLB"
// 	KubeVip              string = "kube-vip"
// 	ServiceLb            string = "servicelb"
// 	IngressAmbassador    string = "ambassador"
// 	IngressAmbassadorAPI string = "ambassadorApi"
// 	IngressHaproxy       string = "aaproxy"
// 	IngressContour       string = "contour"
// 	IngressNginx         string = "nginx"
// 	IngressKing          string = "king"
// 	IngressTraefik       string = "traefik"
// 	Flannel              string = "flannel"
// 	Calico               string = "calico"
// 	Cilium               string = "cilium"
// 	Vxlan                string = "vxlan"
// 	None                 string = "none"
// 	IpSec                string = "ipsec"
// 	HostGw               string = "host-gw"
// 	WireGuard            string = "wireguard"
// 	IpIp                 string = "ipip"
// 	Bgp                  string = "bgp"
// )

// var IngressControllers = []string{IngressAmbassador, IngressAmbassadorAPI, IngressContour, IngressHaproxy, IngressKing, IngressNginx, IngressTraefik}
// var FlannelBackends = []string{Vxlan, None, IpSec, HostGw, WireGuard}
// var CalicoBackends = []string{Vxlan, IpIp, WireGuard, Bgp}
// var CiliumBackends = []string{Vxlan, IpIp, WireGuard}

// var CNIplugins = []string{Flannel, Calico, Cilium}

// DefaultTmpfsMounts specifies tmpfs mounts that are required for all k3d nodes
var DefaultTmpfsMounts = []string{
	"/run",
	"/var/run",
}

// DefaultNodeEnv defines some default environment variables that should be set on every node
var DefaultNodeEnv = []string{
	"K3S_KUBECONFIG_OUTPUT=/output/kubeconfig.yaml",
}

// DefaultEnvPrefix DEPLOY_MYENV
const DefaultEnvPrefix = "DEPLOY"

// DefaultConfigDirName defines the name of the config directory
const DefaultConfigDirName = ".deploy" // should end up in $HOME/

// DefaultSshPort defines the default SSH Port
const DefaultSshPort = 22

// NodeOpts set of options one can connection to node
type NodeOpts struct {
	// Name The bastion Name
	Name string `mapstructure:"name" yaml:"name" json:"name,omitempty"`
	Host string `mapstructure:"host,omitempty" yaml:"host" json:"host"`
	// User SSH user is empty use root
	// +optional
	User string `mapstructure:"user,omitempty" yaml:"user,omitempty" json:"user,omitempty"`
	// Protocol string `mapstructure:"protocol,omitempty" yaml:"protocol,omitempty" json:"protocol,omitempty"` // default: http
	Password string `mapstructure:"password,omitempty" yaml:"password,omitempty" json:"password,omitempty"`
	// SshPort specifies the port the SSH bastion host.
	// Defaults to 22.
	// +optional
	SshPort int32 `mapstructure:"sshPort" yaml:"sshPort" json:"sshPort,omitempty"`
	// SSHAuthorizedKey specifies a list of ssh authorized keys for the user
	// +optional
	SSHAuthorizedKey string `mapstructure:"sshAuthorizedKey" yaml:"sshAuthorizedKey" json:"sshAuthorizedKey,omitempty"`
	// RemoteConnectionString TODO: tranclate строка подключения к удаленному хосту если через bastion
	// +optional
	RemoteConnectionString string `mapstructure:"remoteConnectionString,omitempty" yaml:"remoteConnectionString,omitempty" json:"remoteConnectionString,omitempty"`
	// RemoteSudo TODO: tranclate если через bastion и пользовател на приватном хосте не root устанавливается true
	// +optional
	Workdir    string        `mapstructure:"workdir,omitempty" yaml:"workdir" json:"workdir"`
	RemoteSudo string        `mapstructure:"remoteSudo,omitempty" yaml:"remoteSudo,omitempty" json:"remoteSudo,omitempty"`
	Env        EnvList       `mapstructure:"env,omitempty" yaml:"env" json:"env,omitempty"`
	Bastion    string        `mapstructure:"bastion,omitempty" yaml:"bastion" json:"bastion,omitempty"` // Jump host for the environment
	Wait       bool          `mapstructure:"wait,omitempty" yaml:"wait" json:"wait,omitempty"`
	Timeout    time.Duration `mapstructure:"timeout,omitempty" yaml:"timeout" json:"timeout,omitempty"`
}

// GetWorkdir example: cd /home/ubuntu/app_name;
func (s *NodeOpts) GetWorkdir() (result string) {
	if len(s.Workdir) > 0 {
		result = fmt.Sprintf("cd %s; ", s.Workdir)
	}
	return result
}

// GetEnvExport
func (s *NodeOpts) GetEnvExport() (result string) {
	for _, item := range s.Env {
		result = fmt.Sprintf("%s%s ", result, item.AsExport())
	}
	return result
}

// SetEnv setting EnvList
func (s *NodeOpts) SetEnv(items []string) {
	for _, item := range items {
		row := strings.SplitN(item, "=", 2)
		if len(row) > 1 {
			res := &EnvVar{
				Key:   strings.TrimSpace(row[0]),
				Value: strings.TrimSpace(row[1]),
			}
			s.Env = append(s.Env, res)
			// log.Warnf("key: %s | val: %s", res.Key, res.Value)
		}
	}
}
