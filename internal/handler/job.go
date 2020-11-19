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
	"fmt"

	log "github.com/sirupsen/logrus"
	"github.com/tejabeta/kubectl-ttl/pkg/k8s"
	batchv1 "k8s.io/api/batch/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// CreateJob function is used to create a job to execute ttl
func CreateJob(ns string, kind string, name string, sa string, time int) {
	var ttlSecondsAfterFinished, backOffLimit int32 = 300, 3

	jobContainer := v1.Container{
		Name:    "ttl-" + kind + "-Container",
		Image:   "bitnami/kubectl",
		Command: []string{"kubectl", "delete", kind + "/" + name, "-n", ns},
	}

	initContainer := v1.Container{
		Name:    "ttl-" + kind + "-Init-Container",
		Image:   "busybox:1.28",
		Command: []string{"sh", "-c", "sleep " + fmt.Sprint(time) + "m"},
	}

	job := &batchv1.Job{
		ObjectMeta: metav1.ObjectMeta{Name: "ttl-" + kind + "-job", Namespace: ns},
		Spec: batchv1.JobSpec{
			TTLSecondsAfterFinished: &ttlSecondsAfterFinished,
			BackoffLimit:            &backOffLimit,
			Template: v1.PodTemplateSpec{
				Spec: v1.PodSpec{
					ServiceAccountName: sa,
					Containers:         []v1.Container{jobContainer},
					InitContainers:     []v1.Container{initContainer},
				},
			},
		},
	}

	context, err := getContext()
	if err != nil {
		log.Fatalln(err)
	}

	options, err := k8s.GetOpts(context, ns)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = options.CreateJob(job)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("ttl-", kind, "-job created successfully")
}

// CheckJob lets us validate and identify the roles
func CheckJob(ns string, name string) bool {
	context, err := getContext()
	if err != nil {
		log.Fatalln(err)
	}

	options, err := k8s.GetOpts(context, ns)
	if err != nil {
		log.Fatalln(err)
	}

	results, err := options.GetJob()
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
