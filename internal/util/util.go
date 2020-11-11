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
)

// ResInfo a struct that defines information about resources
type ResInfo struct {
	name      string
	kind      string
	namespace string
}

// ValidResList contains a list of all the valid resources
var ValidResList = map[string]bool{
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

func unique(s []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// IsJSON a function to validate the provided input is json
func IsJSON(s []byte) bool {
	return bytes.HasPrefix(bytes.TrimLeftFunc(s, unicode.IsSpace), []byte{'{'})
}

// IsYAML a function to validate the provided input is json
func IsYAML(s []byte) bool {
	return bytes.HasPrefix(bytes.TrimLeftFunc(s, unicode.IsSpace), []byte{'a', 'p', 'i', 'V', 'e', 'r', 's', 'i', 'o', 'n'})
}
