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
	core_v1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
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
			in.(*Caffe2Job).DeepCopyInto(out.(*Caffe2Job))
			return nil
		}, InType: reflect.TypeOf(&Caffe2Job{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Caffe2JobCondition).DeepCopyInto(out.(*Caffe2JobCondition))
			return nil
		}, InType: reflect.TypeOf(&Caffe2JobCondition{})},
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
			in.(*TerminationPolicySpec).DeepCopyInto(out.(*TerminationPolicySpec))
			return nil
		}, InType: reflect.TypeOf(&TerminationPolicySpec{})},
	)
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
func (in *Caffe2JobCondition) DeepCopyInto(out *Caffe2JobCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Caffe2JobCondition.
func (in *Caffe2JobCondition) DeepCopy() *Caffe2JobCondition {
	if in == nil {
		return nil
	}
	out := new(Caffe2JobCondition)
	in.DeepCopyInto(out)
	return out
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
		*out = make(map[Caffe2ReplicaType]*Caffe2ReplicaSpec, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = new(Caffe2ReplicaSpec)
				val.DeepCopyInto((*out)[key])
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
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Caffe2JobCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ReplicaStatuses != nil {
		in, out := &in.ReplicaStatuses, &out.ReplicaStatuses
		*out = make(map[Caffe2ReplicaType]*Caffe2ReplicaStatus, len(*in))
		for key, val := range *in {
			if val == nil {
				(*out)[key] = nil
			} else {
				(*out)[key] = new(Caffe2ReplicaStatus)
				val.DeepCopyInto((*out)[key])
			}
		}
	}
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.CompletionTime != nil {
		in, out := &in.CompletionTime, &out.CompletionTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Time)
			(*in).DeepCopyInto(*out)
		}
	}
	if in.LastReconcileTime != nil {
		in, out := &in.LastReconcileTime, &out.LastReconcileTime
		if *in == nil {
			*out = nil
		} else {
			*out = new(v1.Time)
			(*in).DeepCopyInto(*out)
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
			*out = new(core_v1.PodTemplateSpec)
			(*in).DeepCopyInto(*out)
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
