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

	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	testclient "k8s.io/client-go/kubernetes/fake"
)

func TestGetJob(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &batchv1.Job{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-job", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := cs.BatchV1().Jobs(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	output, err := options.GetJob()
	if err != nil {
		t.Fatal(err.Error())
	}

	for _, v := range output.Items {
		if v.Name != "unit-test-job" {
			t.Errorf("Error while getting job")
		}
	}

}

func TestDeleteJob(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &batchv1.Job{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-job", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := cs.BatchV1().Jobs(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.DeleteJob("unit-test-job")
	if err != nil {
		t.Fatal(err.Error())
	}

	err = options.DeleteJob("unit-test-job-1")
	if err == nil {
		t.Errorf("Error while deleting non-existence job")
	}

}

func TestCreateJob(t *testing.T) {

	cs := testclient.NewSimpleClientset()
	input := &batchv1.Job{ObjectMeta: metav1.ObjectMeta{Name: "unit-test-job", ResourceVersion: "12345"}}

	options := Options{
		clientset: cs,
		namespace: "unit-test-namespace",
	}

	_, err := options.CreateJob(input)
	if err != nil {
		t.Fatal(err.Error())
	}

	output, err := cs.BatchV1().Jobs(options.namespace).Get(context.TODO(), "unit-test-job", metav1.GetOptions{})
	if err != nil {
		t.Fatal(err.Error())
	}

	if output.Name != "unit-test-job" {
		t.Errorf("Error while retreiving created job")
	}

	_, err = cs.BatchV1().Jobs(options.namespace).Create(context.TODO(), input, metav1.CreateOptions{})
	if err == nil {
		t.Errorf("Error while trying to create a duplicate job")
	}

}
