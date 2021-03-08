// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsTrait) DeepCopyInto(out *MetricsTrait) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsTrait.
func (in *MetricsTrait) DeepCopy() *MetricsTrait {
	if in == nil {
		return nil
	}
	out := new(MetricsTrait)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetricsTrait) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsTraitList) DeepCopyInto(out *MetricsTraitList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MetricsTrait, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsTraitList.
func (in *MetricsTraitList) DeepCopy() *MetricsTraitList {
	if in == nil {
		return nil
	}
	out := new(MetricsTraitList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MetricsTraitList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsTraitSpec) DeepCopyInto(out *MetricsTraitSpec) {
	*out = *in
	in.ScrapeService.DeepCopyInto(&out.ScrapeService)
	out.WorkloadReference = in.WorkloadReference
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsTraitSpec.
func (in *MetricsTraitSpec) DeepCopy() *MetricsTraitSpec {
	if in == nil {
		return nil
	}
	out := new(MetricsTraitSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetricsTraitStatus) DeepCopyInto(out *MetricsTraitStatus) {
	*out = *in
	in.ConditionedStatus.DeepCopyInto(&out.ConditionedStatus)
	out.Port = in.Port
	if in.SelectorLabels != nil {
		in, out := &in.SelectorLabels, &out.SelectorLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetricsTraitStatus.
func (in *MetricsTraitStatus) DeepCopy() *MetricsTraitStatus {
	if in == nil {
		return nil
	}
	out := new(MetricsTraitStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScapeServiceEndPoint) DeepCopyInto(out *ScapeServiceEndPoint) {
	*out = *in
	out.TargetPort = in.TargetPort
	if in.TargetSelector != nil {
		in, out := &in.TargetSelector, &out.TargetSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScapeServiceEndPoint.
func (in *ScapeServiceEndPoint) DeepCopy() *ScapeServiceEndPoint {
	if in == nil {
		return nil
	}
	out := new(ScapeServiceEndPoint)
	in.DeepCopyInto(out)
	return out
}
