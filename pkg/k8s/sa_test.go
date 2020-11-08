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

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	testclient "k8s.io/client-go/kubernetes/fake"
)

func TestGetSA(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &v1.ServiceAccount{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-sa", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := cs.CoreV1().ServiceAccounts(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	output, err := options.GetSA()
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, v := range output.Items {
		if v.Name != "unit-test-sa" {
			t.Errorf("Error while getting service accounts")
		}
	}

}

func TestDeleteSA(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &v1.ServiceAccount{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-sa", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := cs.CoreV1().ServiceAccounts(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.DeleteSA("unit-test-sa")
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.DeleteRBinding("unit-test-sa-1")
	if err == nil {
		t.Errorf("Error while deleting a non existence service account")
	}

}

func TestCreateSA(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &v1.ServiceAccount{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-sa", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := options.CreateSA(input)
	if err != nil {
		t.Fatal(err.Error())
	}

	output, err := cs.CoreV1().ServiceAccounts(options.namespace).Get(context.TODO(), "unit-test-sa", metav1.GetOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	if output.Name != "unit-test-sa" {
		t.Errorf("Created service account doesn't exist")
	}

	_, err = cs.CoreV1().ServiceAccounts(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err == nil {
		t.Errorf("Error while creating a duplicate service account")
	}

}
