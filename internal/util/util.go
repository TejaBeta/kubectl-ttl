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
	"bytes"
	"unicode"

	log "github.com/sirupsen/logrus"
	"github.com/tidwall/gjson"
)

var resourceList = map[string]bool{
	"Pod":                   true,
	"Service":               true,
	"Ingress":               true,
	"ConfigMap":             true,
	"Secret":                true,
	"ReplicaSet":            true,
	"Deployment":            true,
	"PersistentVolumeClaim": true,
	"PersistentVolume":      true,
	"ServiceAccount":        true,
}

// IsJSON a function to validate the provided input is json
func IsJSON(s []byte) bool {
	return bytes.HasPrefix(bytes.TrimLeftFunc(s, unicode.IsSpace), []byte{'{'})
}

// IsYAML a function to validate the provided input is json
func IsYAML(s []byte) bool {
	return bytes.HasPrefix(bytes.TrimLeftFunc(s, unicode.IsSpace), []byte{'a', 'p', 'i', 'V', 'e', 'r', 's', 'i', 'o', 'n'})
}

// IsResValid a function to parse and validate if the input resouces are valid for ttl
func IsResValid(s string) bool {
	if !gjson.Valid(s) {
		log.Fatal("Invalid input")
	}

	if gjson.Get(s, "kind").String() == "List" {
		result := gjson.Get(s, "items.#.kind")
		for _, data := range result.Array() {
			if !resourceList[data.String()] {
				return false
			}
		}
	} else {
		if !(resourceList[gjson.Get(s, "kind").String()]) {
			return false
		}
	}
	return true
}