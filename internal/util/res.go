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

// ResType a function to parse and get all the resource type from the input
func ResType(s string) []string {
	var output []string
	if !gjson.Valid(s) {
		log.Fatal("Invalid input")
	}

	if gjson.Get(s, "kind").String() == "List" {
		result := gjson.Get(s, "items.#.kind")
		for _, data := range result.Array() {
			output = append(output, data.String())
		}
	} else {
		output = append(output, gjson.Get(s, "kind").String())
	}
	return output
}
