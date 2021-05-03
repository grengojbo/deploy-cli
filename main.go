/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	// "flag"
	// "os"

	"github.com/grengojbo/deploy-cli/cmd"
)

var (
// scheme   = runtime.NewScheme()
// setupLog = ctrl.Log.WithName("setup")
)

// func init() {
// 	_ = clientgoscheme.AddToScheme(scheme)

// 	_ = k3sv1alpha1.AddToScheme(scheme)
// 	// +kubebuilder:scaffold:scheme
// }

func main() {

	cmd.Execute()
	// if err := cmd.Execute(); err != nil {
	// 	os.Exit(1)
	// }
}
