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

package util

import (
	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

// GetResDetails returns information related to resources
func GetResDetails(s string) []ResInfo {
	if !gjson.Valid(s) {
		log.Fatal("Invalid input")
	}

	var output []ResInfo

	if gjson.Get(s, "kind").String() == "List" {
		kinds := gjson.Get(s, "items.#.kind").Array()
		names := gjson.Get(s, "items.#.metadata.name").Array()
		namespaces := gjson.Get(s, "items.#.metadata.namespace").Array()
		for i, v := range kinds {
			value := ResInfo{name: names[i].String(), kind: v.String(), namespace: namespaces[i].String()}
			output = append(output, value)
		}
	}

	return output
}
