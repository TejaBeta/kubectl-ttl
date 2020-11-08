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

	batchv1 "k8s.io/api/batch/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetJob returns all the jobs in the given namespace and clientset
func (kOpts *Options) GetJob() (result *batchv1.JobList, err error) {
	result, err = kOpts.clientset.
		BatchV1().
		Jobs(kOpts.namespace).
		List(context.TODO(), metav1.ListOptions{})
	return
}

// DeleteJob method to delete a job with the name
func (kOpts *Options) DeleteJob(name string) (err error) {
	err = kOpts.clientset.
		BatchV1().
		Jobs(kOpts.namespace).
		Delete(context.TODO(), name, metav1.DeleteOptions{})
	return
}

// CreateJob method to create a pvc
func (kOpts *Options) CreateJob(job *batchv1.Job) (result *batchv1.Job, err error) {
	result, err = kOpts.clientset.
		BatchV1().
		Jobs(kOpts.namespace).
		Create(context.TODO(), job, metav1.CreateOptions{})
	return
}
