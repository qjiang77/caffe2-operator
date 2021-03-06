// Copyright 2018 The Kubeflow Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v1alpha1

import (
	"fmt"
	"time"

	"github.com/golang/protobuf/proto"
	"k8s.io/apimachinery/pkg/runtime"
)

func addDefaultingFuncs(scheme *runtime.Scheme) error {
	return RegisterDefaults(scheme)
}

// SetDefaults_Caffe2Job sets any unspecified values to defaults
func SetDefaults_Caffe2Job(obj *Caffe2Job) {
	c := &obj.Spec

	if c.RuntimeID == "" {
		c.RuntimeID = fmt.Sprintf("%d", time.Now().Unix())
	}
	// Check that each replica has a Caffe2 container.
	r := c.ReplicaSpecs
	/*
		if r.Caffe2Port == nil {
			r.Caffe2Port = proto.Int32(Caffe2Port)
		}
	*/
	if r.Template.Spec.RestartPolicy == "" {
		r.Template.Spec.RestartPolicy = "Never"
	}

	if r.Replicas == nil {
		r.Replicas = proto.Int32(1)
	}
	if c.Backend == "" {
		if *r.Replicas == 1 {
			c.Backend = NoneBackendType
		} else {
			c.Backend = RedisBackendType
		}
	}

	if c.TerminationPolicy == nil {
		c.TerminationPolicy = &TerminationPolicySpec{
			Chief: &ChiefSpec{
				ReplicaName:  "WORKER",
				ReplicaIndex: 0,
			},
		}
	}
}
