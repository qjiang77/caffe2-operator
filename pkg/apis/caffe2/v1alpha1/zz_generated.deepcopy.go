// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AcceleratorConfig).DeepCopyInto(out.(*AcceleratorConfig))
			return nil
		}, InType: reflect.TypeOf(&AcceleratorConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*AcceleratorVolume).DeepCopyInto(out.(*AcceleratorVolume))
			return nil
		}, InType: reflect.TypeOf(&AcceleratorVolume{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Caffe2Job).DeepCopyInto(out.(*Caffe2Job))
			return nil
		}, InType: reflect.TypeOf(&Caffe2Job{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Caffe2JobList).DeepCopyInto(out.(*Caffe2JobList))
			return nil
		}, InType: reflect.TypeOf(&Caffe2JobList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Caffe2JobSpec).DeepCopyInto(out.(*Caffe2JobSpec))
			return nil
		}, InType: reflect.TypeOf(&Caffe2JobSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Caffe2JobStatus).DeepCopyInto(out.(*Caffe2JobStatus))
			return nil
		}, InType: reflect.TypeOf(&Caffe2JobStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Caffe2ReplicaSpec).DeepCopyInto(out.(*Caffe2ReplicaSpec))
			return nil
		}, InType: reflect.TypeOf(&Caffe2ReplicaSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Caffe2ReplicaStatus).DeepCopyInto(out.(*Caffe2ReplicaStatus))
			return nil
		}, InType: reflect.TypeOf(&Caffe2ReplicaStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ChiefSpec).DeepCopyInto(out.(*ChiefSpec))
			return nil
		}, InType: reflect.TypeOf(&ChiefSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*ControllerConfig).DeepCopyInto(out.(*ControllerConfig))
			return nil
		}, InType: reflect.TypeOf(&ControllerConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*EnvironmentVariableConfig).DeepCopyInto(out.(*EnvironmentVariableConfig))
			return nil
		}, InType: reflect.TypeOf(&EnvironmentVariableConfig{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*TerminationPolicySpec).DeepCopyInto(out.(*TerminationPolicySpec))
			return nil
		}, InType: reflect.TypeOf(&TerminationPolicySpec{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceleratorConfig) DeepCopyInto(out *AcceleratorConfig) {
	*out = *in
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]AcceleratorVolume, len(*in))
		copy(*out, *in)
	}
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make([]EnvironmentVariableConfig, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceleratorConfig.
func (in *AcceleratorConfig) DeepCopy() *AcceleratorConfig {
	if in == nil {
		return nil
	}
	out := new(AcceleratorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AcceleratorVolume) DeepCopyInto(out *AcceleratorVolume) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AcceleratorVolume.
func (in *AcceleratorVolume) DeepCopy() *AcceleratorVolume {
	if in == nil {
		return nil
	}
	out := new(AcceleratorVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Caffe2Job) DeepCopyInto(out *Caffe2Job) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Caffe2Job.
func (in *Caffe2Job) DeepCopy() *Caffe2Job {
	if in == nil {
		return nil
	}
	out := new(Caffe2Job)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Caffe2Job) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Caffe2JobList) DeepCopyInto(out *Caffe2JobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Caffe2Job, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Caffe2JobList.
func (in *Caffe2JobList) DeepCopy() *Caffe2JobList {
	if in == nil {
		return nil
	}
	out := new(Caffe2JobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Caffe2JobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Caffe2JobSpec) DeepCopyInto(out *Caffe2JobSpec) {
	*out = *in
	if in.ReplicaSpecs != nil {
		in, out := &in.ReplicaSpecs, &out.ReplicaSpecs
		*out = make([]*Caffe2ReplicaSpec, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(Caffe2ReplicaSpec)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	if in.TerminationPolicy != nil {
		in, out := &in.TerminationPolicy, &out.TerminationPolicy
		if *in == nil {
			*out = nil
		} else {
			*out = new(TerminationPolicySpec)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Caffe2JobSpec.
func (in *Caffe2JobSpec) DeepCopy() *Caffe2JobSpec {
	if in == nil {
		return nil
	}
	out := new(Caffe2JobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Caffe2JobStatus) DeepCopyInto(out *Caffe2JobStatus) {
	*out = *in
	if in.ReplicaStatuses != nil {
		in, out := &in.ReplicaStatuses, &out.ReplicaStatuses
		*out = make([]*Caffe2ReplicaStatus, len(*in))
		for i := range *in {
			if (*in)[i] == nil {
				(*out)[i] = nil
			} else {
				(*out)[i] = new(Caffe2ReplicaStatus)
				(*in)[i].DeepCopyInto((*out)[i])
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Caffe2JobStatus.
func (in *Caffe2JobStatus) DeepCopy() *Caffe2JobStatus {
	if in == nil {
		return nil
	}
	out := new(Caffe2JobStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Caffe2ReplicaSpec) DeepCopyInto(out *Caffe2ReplicaSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	if in.Template != nil {
		in, out := &in.Template, &out.Template
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.PodTemplateSpec)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.Caffe2Port != nil {
		in, out := &in.Caffe2Port, &out.Caffe2Port
		if *in == nil {
			*out = nil
		} else {
			*out = new(int32)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Caffe2ReplicaSpec.
func (in *Caffe2ReplicaSpec) DeepCopy() *Caffe2ReplicaSpec {
	if in == nil {
		return nil
	}
	out := new(Caffe2ReplicaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Caffe2ReplicaStatus) DeepCopyInto(out *Caffe2ReplicaStatus) {
	*out = *in
	if in.ReplicasStates != nil {
		in, out := &in.ReplicasStates, &out.ReplicasStates
		*out = make(map[ReplicaState]int, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Caffe2ReplicaStatus.
func (in *Caffe2ReplicaStatus) DeepCopy() *Caffe2ReplicaStatus {
	if in == nil {
		return nil
	}
	out := new(Caffe2ReplicaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChiefSpec) DeepCopyInto(out *ChiefSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChiefSpec.
func (in *ChiefSpec) DeepCopy() *ChiefSpec {
	if in == nil {
		return nil
	}
	out := new(ChiefSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerConfig) DeepCopyInto(out *ControllerConfig) {
	*out = *in
	if in.Accelerators != nil {
		in, out := &in.Accelerators, &out.Accelerators
		*out = make(map[string]AcceleratorConfig, len(*in))
		for key, val := range *in {
			newVal := new(AcceleratorConfig)
			val.DeepCopyInto(newVal)
			(*out)[key] = *newVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerConfig.
func (in *ControllerConfig) DeepCopy() *ControllerConfig {
	if in == nil {
		return nil
	}
	out := new(ControllerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvironmentVariableConfig) DeepCopyInto(out *EnvironmentVariableConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvironmentVariableConfig.
func (in *EnvironmentVariableConfig) DeepCopy() *EnvironmentVariableConfig {
	if in == nil {
		return nil
	}
	out := new(EnvironmentVariableConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TerminationPolicySpec) DeepCopyInto(out *TerminationPolicySpec) {
	*out = *in
	if in.Chief != nil {
		in, out := &in.Chief, &out.Chief
		if *in == nil {
			*out = nil
		} else {
			*out = new(ChiefSpec)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TerminationPolicySpec.
func (in *TerminationPolicySpec) DeepCopy() *TerminationPolicySpec {
	if in == nil {
		return nil
	}
	out := new(TerminationPolicySpec)
	in.DeepCopyInto(out)
	return out
}
