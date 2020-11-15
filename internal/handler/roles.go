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
	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateRole is a function that creates roles for ttl to work
func CreateRole(ns string, kind string, resName string) {
	rule := &rbacv1.PolicyRule{
		Verbs:         []string{"delete"},
		APIGroups:     []string{"*"},
		ResourceNames: []string{resName},
		Resources:     []string{kind},
	}

	role := &rbacv1.Role{
		ObjectMeta: metav1.ObjectMeta{Name: "ttl-" + kind + "-role", Namespace: ns},
		Rules:      []rbacv1.PolicyRule{*rule},
	}

	context, err := getContext()
	if err != nil {
		log.Fatalln(err)
	}

	options, err := k8s.GetOpts(context, ns)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = options.CreateRole(role)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Role ttl-", kind, "-role is created in namespace ", ns)
}