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

package ttl

import (
	log "github.com/sirupsen/logrus"
	ttlOpts "github.com/tejabeta/kubectl-ttl/internal/options"
)

// KubectlTTL is the main function that acts as the entry point
func KubectlTTL(options *ttlOpts.Options) {
	log.Println(options.AllResources, options.Namespace, options.TimeToLive)
}
