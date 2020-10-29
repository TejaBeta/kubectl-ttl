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

package ttl

import (
	log "github.com/sirupsen/logrus"
	"github.com/tejabeta/kopy/pkg/koperator"
	ttlOpts "github.com/tejabeta/kubectl-ttl/internal/options"
	appv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1beta1 "k8s.io/api/extensions/v1beta1"
	rbacv1 "k8s.io/api/rbac/v1"
)

// KopyResources as name suggests a struct type to hold all the resources
type ttlResources struct {
	Deployments  *[]appv1.Deployment
	ConfigMaps   *[]corev1.ConfigMap
	Roles        *[]rbacv1.Role
	RoleBindings *[]rbacv1.RoleBinding
	Secrets      *[]corev1.Secret
	Services     *[]corev1.Service
	Ingresses    *[]v1beta1.Ingress
}

// KubectlTTL is the main function that acts as the entry point
func KubectlTTL(options *ttlOpts.Options) {
	kOpts, err := koperator.GetOpts(options.Context, options.Namespace)
	if err != nil {
		log.Errorln(err)
		return
	}

	if isNS(kOpts) {
		log.Println("Namespace ", options.Namespace, " exits")
	} else {
		log.Error("No namespace ", options.Namespace, " found")
	}

	return
}

func isNS(kOpts *koperator.Options) bool {
	_, err := kOpts.GetNS()
	if err != nil {
		return false
	}
	return true
}
