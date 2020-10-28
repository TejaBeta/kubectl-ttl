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
package options

import (
	"github.com/tejabeta/kubectl-ttl/internal/context"
	"k8s.io/client-go/rest"
)

// Options is struct holding all the options for the tool to work
type Options struct {
	Namespace   string
	AllResource bool
	Context     *rest.Config
}

// GetOptions is a function helps to retreive certain options
func GetOptions() (*Options, error) {
	return context.GetContext()
}
