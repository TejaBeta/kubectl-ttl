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

package handler

import (
	log "github.com/sirupsen/logrus"
	"github.com/tejabeta/kubectl-ttl/pkg/k8s"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateSA helps to create a service account
func CreateSA(ns string) {
	sa := &v1.ServiceAccount{ObjectMeta: metav1.ObjectMeta{Name: "ttl-sa", Namespace: ns}}

	context, err := getContext()
	if err != nil {
		log.Fatalln(err)
	}

	options, err := k8s.GetOpts(context, ns)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = options.CreateSA(sa)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Service Account ttl-sa is created in namespace ", ns)
}

// CheckSA validates and identifies if a service account already exists
func CheckSA(ns string, name string) bool {
	context, err := getContext()
	if err != nil {
		log.Fatalln(err)
	}

	options, err := k8s.GetOpts(context, ns)
	if err != nil {
		log.Fatalln(err)
	}

	results, err := options.GetSA()
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range results.Items {
		if v.Name == name {
			return true
		}
	}

	return false
}
