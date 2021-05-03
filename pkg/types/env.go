package types

import "fmt"

// EnvVar represents an environment variable
type EnvVar struct {
	Key   string
	Value string
}

func (e EnvVar) String() string {
	return e.Key + `=` + e.Value
}

// AsExport returns the environment variable as a bash export statement
func (e EnvVar) AsExport() string {
	return `export ` + e.Key + `="` + e.Value + `";`
}

// EnvList is a list of environment variables that maps to a YAML map,
// but maintains order, enabling late variables to reference early variables.
type EnvList []*EnvVar

func (e EnvList) Slice() []string {
	envs := make([]string, len(e))
	for i, env := range e {
		envs[i] = env.String()
	}
	return envs
}

// GetEnvExport example export bbb=sss; export aaa=zzz; export qqq=www;
// ./deploy-cli --set="bbb=sss,aaa=zzz" --set=qqq=www run
func GetEnvExport(el []string) (result string) {
	for _, s := range el {
		result = fmt.Sprintf("%sexport %s; ", result, s)
	}
	return result
}
