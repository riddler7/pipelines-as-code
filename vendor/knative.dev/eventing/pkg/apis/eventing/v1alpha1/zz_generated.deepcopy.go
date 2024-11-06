//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021 The Knative Authors

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	eventingv1 "knative.dev/eventing/pkg/apis/eventing/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventPolicy) DeepCopyInto(out *EventPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventPolicy.
func (in *EventPolicy) DeepCopy() *EventPolicy {
	if in == nil {
		return nil
	}
	out := new(EventPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventPolicyFromReference) DeepCopyInto(out *EventPolicyFromReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventPolicyFromReference.
func (in *EventPolicyFromReference) DeepCopy() *EventPolicyFromReference {
	if in == nil {
		return nil
	}
	out := new(EventPolicyFromReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventPolicyList) DeepCopyInto(out *EventPolicyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]EventPolicy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventPolicyList.
func (in *EventPolicyList) DeepCopy() *EventPolicyList {
	if in == nil {
		return nil
	}
	out := new(EventPolicyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *EventPolicyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventPolicySelector) DeepCopyInto(out *EventPolicySelector) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(v1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	if in.TypeMeta != nil {
		in, out := &in.TypeMeta, &out.TypeMeta
		*out = new(v1.TypeMeta)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventPolicySelector.
func (in *EventPolicySelector) DeepCopy() *EventPolicySelector {
	if in == nil {
		return nil
	}
	out := new(EventPolicySelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventPolicySpec) DeepCopyInto(out *EventPolicySpec) {
	*out = *in
	if in.To != nil {
		in, out := &in.To, &out.To
		*out = make([]EventPolicySpecTo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = make([]EventPolicySpecFrom, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Filters != nil {
		in, out := &in.Filters, &out.Filters
		*out = make([]eventingv1.SubscriptionsAPIFilter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventPolicySpec.
func (in *EventPolicySpec) DeepCopy() *EventPolicySpec {
	if in == nil {
		return nil
	}
	out := new(EventPolicySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventPolicySpecFrom) DeepCopyInto(out *EventPolicySpecFrom) {
	*out = *in
	if in.Ref != nil {
		in, out := &in.Ref, &out.Ref
		*out = new(EventPolicyFromReference)
		**out = **in
	}
	if in.Sub != nil {
		in, out := &in.Sub, &out.Sub
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventPolicySpecFrom.
func (in *EventPolicySpecFrom) DeepCopy() *EventPolicySpecFrom {
	if in == nil {
		return nil
	}
	out := new(EventPolicySpecFrom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventPolicySpecTo) DeepCopyInto(out *EventPolicySpecTo) {
	*out = *in
	if in.Ref != nil {
		in, out := &in.Ref, &out.Ref
		*out = new(EventPolicyToReference)
		**out = **in
	}
	if in.Selector != nil {
		in, out := &in.Selector, &out.Selector
		*out = new(EventPolicySelector)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventPolicySpecTo.
func (in *EventPolicySpecTo) DeepCopy() *EventPolicySpecTo {
	if in == nil {
		return nil
	}
	out := new(EventPolicySpecTo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventPolicyStatus) DeepCopyInto(out *EventPolicyStatus) {
	*out = *in
	in.Status.DeepCopyInto(&out.Status)
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventPolicyStatus.
func (in *EventPolicyStatus) DeepCopy() *EventPolicyStatus {
	if in == nil {
		return nil
	}
	out := new(EventPolicyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EventPolicyToReference) DeepCopyInto(out *EventPolicyToReference) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EventPolicyToReference.
func (in *EventPolicyToReference) DeepCopy() *EventPolicyToReference {
	if in == nil {
		return nil
	}
	out := new(EventPolicyToReference)
	in.DeepCopyInto(out)
	return out
}
