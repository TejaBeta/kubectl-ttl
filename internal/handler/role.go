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

// CreateRB is a function to create role binding to the above role
func CreateRB(ns string, sa string, role string) {

	subject := rbacv1.Subject{
		Kind:      "ServiceAccount",
		Name:      sa,
		Namespace: ns,
	}

	roleRef := rbacv1.RoleRef{
		APIGroup: "rbac.authorization.k8s.io",
		Kind:     "Role",
		Name:     role,
	}

	rb := &rbacv1.RoleBinding{
		ObjectMeta: metav1.ObjectMeta{Name: "ttl-rolebinding", Namespace: ns},
		Subjects:   []rbacv1.Subject{subject},
		RoleRef:    roleRef,
	}

	context, err := getContext()
	if err != nil {
		log.Fatalln(err)
	}

	options, err := k8s.GetOpts(context, ns)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = options.CreateRBinding(rb)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Role ttl-rolebinding is created in namespace ", ns)
}

// CheckRole lets us validate and identify the roles
func CheckRole(ns string, name string) bool {
	context, err := getContext()
	if err != nil {
		log.Fatalln(err)
	}

	options, err := k8s.GetOpts(context, ns)
	if err != nil {
		log.Fatalln(err)
	}

	results, err := options.GetRoles()
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

// CheckRB lets us validate and identify the roles
func CheckRB(ns string, name string) bool {
	context, err := getContext()
	if err != nil {
		log.Fatalln(err)
	}

	options, err := k8s.GetOpts(context, ns)
	if err != nil {
		log.Fatalln(err)
	}

	results, err := options.GetRoleBindings()
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
