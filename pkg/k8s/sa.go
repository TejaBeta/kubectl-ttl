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

	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// GetSA returns all the Service accounts in the given namespace and clientset
func (kOpts *Options) GetSA() (result *v1.ServiceAccountList, err error) {
	result, err = kOpts.clientset.
		CoreV1().
		ServiceAccounts(kOpts.namespace).
		List(context.TODO(), metav1.ListOptions{})
	return
}

// DeleteSA method to delete a sa with the name
func (kOpts *Options) DeleteSA(name string) (err error) {
	err = kOpts.clientset.
		CoreV1().
		ServiceAccounts(kOpts.namespace).
		Delete(context.TODO(), name, metav1.DeleteOptions{})
	return
}

// CreateSA method to create a SA
func (kOpts *Options) CreateSA(sa *v1.ServiceAccount) (result *v1.ServiceAccount, err error) {
	result, err = kOpts.clientset.
		CoreV1().
		ServiceAccounts(kOpts.namespace).
		Create(context.TODO(), sa, metav1.CreateOptions{})
	return
}
