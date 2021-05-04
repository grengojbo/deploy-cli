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
package version

// "os"

// "github.com/heroku/docker-registry-client/registry"
// log "github.com/sirupsen/logrus"

// Version is the string that contains version
var Version string

var GitCommit string

// HelperVersionOverride decouples the k3d helper image versions from the main version, if needed
var HelperVersionOverride string

// GetVersion returns the version for cli, it gets it from "git describe --tags" or returns "dev" when doing simple go build
func GetVersion() string {
	if len(Version) == 0 {
		return "v1-dev"
	}
	return Version
}

// GetHelperImageVersion returns the CLI version or 'latest'
// func GetHelperImageVersion() string {
// 	if tag := os.Getenv("K3D_HELPER_IMAGE_TAG"); tag != "" {
// 		log.Infoln("Helper image tag set from env var")
// 		return tag
// 	}
// 	if len(HelperVersionOverride) > 0 {
// 		return HelperVersionOverride
// 	}
// 	if len(Version) == 0 {
// 		return "latest"
// 	}
// 	return Version
// }
