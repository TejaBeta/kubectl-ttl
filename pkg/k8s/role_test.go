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

package k8s

import (
	"context"
	"testing"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	testclient "k8s.io/client-go/kubernetes/fake"
)

func TestGetRoles(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &rbacv1.Role{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-role", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := cs.RbacV1().Roles(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	output, err := options.GetRoles()
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, v := range output.Items {
		if v.Name != "unit-test-role" {
			t.Errorf("Error while getting role")
		}
	}

}

func TestDeleteRole(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &rbacv1.Role{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-role", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := cs.RbacV1().Roles(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.DeleteRole("unit-test-role")
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.DeleteRole("unit-test-role-1")
	if err == nil {
		t.Errorf("Error while deleting non-existence role")
	}
}

func TestCreateRole(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &rbacv1.Role{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-role", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := options.CreateRole(input)
	if err != nil {
		t.Fatal(err.Error())
	}

	output, err := cs.RbacV1().Roles(options.namespace).Get(context.TODO(), "unit-test-role", metav1.GetOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	if output.Name != "unit-test-role" {
		t.Errorf("Error while retreiving created role")
	}

	_, err = cs.RbacV1().Roles(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err == nil {
		t.Errorf("Error while creating duplicate role")
	}

}
