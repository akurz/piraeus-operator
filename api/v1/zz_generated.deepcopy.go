//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterReference) DeepCopyInto(out *ClusterReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterReference.
func (in *ClusterReference) DeepCopy() *ClusterReference {
	if in == nil {
		return nil
	}
	out := new(ClusterReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinstorCluster) DeepCopyInto(out *LinstorCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinstorCluster.
func (in *LinstorCluster) DeepCopy() *LinstorCluster {
	if in == nil {
		return nil
	}
	out := new(LinstorCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LinstorCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinstorClusterList) DeepCopyInto(out *LinstorClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LinstorCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinstorClusterList.
func (in *LinstorClusterList) DeepCopy() *LinstorClusterList {
	if in == nil {
		return nil
	}
	out := new(LinstorClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LinstorClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinstorClusterSpec) DeepCopyInto(out *LinstorClusterSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make([]LinstorControllerProperty, len(*in))
		copy(*out, *in)
	}
	if in.Patches != nil {
		in, out := &in.Patches, &out.Patches
		*out = make([]Patch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinstorClusterSpec.
func (in *LinstorClusterSpec) DeepCopy() *LinstorClusterSpec {
	if in == nil {
		return nil
	}
	out := new(LinstorClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinstorClusterStatus) DeepCopyInto(out *LinstorClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinstorClusterStatus.
func (in *LinstorClusterStatus) DeepCopy() *LinstorClusterStatus {
	if in == nil {
		return nil
	}
	out := new(LinstorClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinstorControllerProperty) DeepCopyInto(out *LinstorControllerProperty) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinstorControllerProperty.
func (in *LinstorControllerProperty) DeepCopy() *LinstorControllerProperty {
	if in == nil {
		return nil
	}
	out := new(LinstorControllerProperty)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinstorNodeProperty) DeepCopyInto(out *LinstorNodeProperty) {
	*out = *in
	if in.ValueFrom != nil {
		in, out := &in.ValueFrom, &out.ValueFrom
		*out = new(LinstorNodePropertyValueFrom)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinstorNodeProperty.
func (in *LinstorNodeProperty) DeepCopy() *LinstorNodeProperty {
	if in == nil {
		return nil
	}
	out := new(LinstorNodeProperty)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinstorNodePropertyValueFrom) DeepCopyInto(out *LinstorNodePropertyValueFrom) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinstorNodePropertyValueFrom.
func (in *LinstorNodePropertyValueFrom) DeepCopy() *LinstorNodePropertyValueFrom {
	if in == nil {
		return nil
	}
	out := new(LinstorNodePropertyValueFrom)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinstorSatellite) DeepCopyInto(out *LinstorSatellite) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinstorSatellite.
func (in *LinstorSatellite) DeepCopy() *LinstorSatellite {
	if in == nil {
		return nil
	}
	out := new(LinstorSatellite)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LinstorSatellite) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinstorSatelliteConfiguration) DeepCopyInto(out *LinstorSatelliteConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinstorSatelliteConfiguration.
func (in *LinstorSatelliteConfiguration) DeepCopy() *LinstorSatelliteConfiguration {
	if in == nil {
		return nil
	}
	out := new(LinstorSatelliteConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LinstorSatelliteConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinstorSatelliteConfigurationList) DeepCopyInto(out *LinstorSatelliteConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LinstorSatelliteConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinstorSatelliteConfigurationList.
func (in *LinstorSatelliteConfigurationList) DeepCopy() *LinstorSatelliteConfigurationList {
	if in == nil {
		return nil
	}
	out := new(LinstorSatelliteConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LinstorSatelliteConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinstorSatelliteConfigurationSpec) DeepCopyInto(out *LinstorSatelliteConfigurationSpec) {
	*out = *in
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Patches != nil {
		in, out := &in.Patches, &out.Patches
		*out = make([]Patch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StoragePools != nil {
		in, out := &in.StoragePools, &out.StoragePools
		*out = make([]LinstorStoragePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make([]LinstorNodeProperty, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinstorSatelliteConfigurationSpec.
func (in *LinstorSatelliteConfigurationSpec) DeepCopy() *LinstorSatelliteConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(LinstorSatelliteConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinstorSatelliteConfigurationStatus) DeepCopyInto(out *LinstorSatelliteConfigurationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinstorSatelliteConfigurationStatus.
func (in *LinstorSatelliteConfigurationStatus) DeepCopy() *LinstorSatelliteConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(LinstorSatelliteConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinstorSatelliteList) DeepCopyInto(out *LinstorSatelliteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LinstorSatellite, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinstorSatelliteList.
func (in *LinstorSatelliteList) DeepCopy() *LinstorSatelliteList {
	if in == nil {
		return nil
	}
	out := new(LinstorSatelliteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LinstorSatelliteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinstorSatelliteSpec) DeepCopyInto(out *LinstorSatelliteSpec) {
	*out = *in
	out.ClusterRef = in.ClusterRef
	if in.Patches != nil {
		in, out := &in.Patches, &out.Patches
		*out = make([]Patch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.StoragePools != nil {
		in, out := &in.StoragePools, &out.StoragePools
		*out = make([]LinstorStoragePool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make([]LinstorNodeProperty, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinstorSatelliteSpec.
func (in *LinstorSatelliteSpec) DeepCopy() *LinstorSatelliteSpec {
	if in == nil {
		return nil
	}
	out := new(LinstorSatelliteSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinstorSatelliteStatus) DeepCopyInto(out *LinstorSatelliteStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinstorSatelliteStatus.
func (in *LinstorSatelliteStatus) DeepCopy() *LinstorSatelliteStatus {
	if in == nil {
		return nil
	}
	out := new(LinstorSatelliteStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinstorStoragePool) DeepCopyInto(out *LinstorStoragePool) {
	*out = *in
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make([]LinstorNodeProperty, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Lvm != nil {
		in, out := &in.Lvm, &out.Lvm
		*out = new(LinstorStoragePoolLvm)
		**out = **in
	}
	if in.LvmThin != nil {
		in, out := &in.LvmThin, &out.LvmThin
		*out = new(LinstorStoragePoolLvmThin)
		**out = **in
	}
	if in.Source != nil {
		in, out := &in.Source, &out.Source
		*out = new(LinstorStoragePoolSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinstorStoragePool.
func (in *LinstorStoragePool) DeepCopy() *LinstorStoragePool {
	if in == nil {
		return nil
	}
	out := new(LinstorStoragePool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinstorStoragePoolLvm) DeepCopyInto(out *LinstorStoragePoolLvm) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinstorStoragePoolLvm.
func (in *LinstorStoragePoolLvm) DeepCopy() *LinstorStoragePoolLvm {
	if in == nil {
		return nil
	}
	out := new(LinstorStoragePoolLvm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinstorStoragePoolLvmThin) DeepCopyInto(out *LinstorStoragePoolLvmThin) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinstorStoragePoolLvmThin.
func (in *LinstorStoragePoolLvmThin) DeepCopy() *LinstorStoragePoolLvmThin {
	if in == nil {
		return nil
	}
	out := new(LinstorStoragePoolLvmThin)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LinstorStoragePoolSource) DeepCopyInto(out *LinstorStoragePoolSource) {
	*out = *in
	if in.HostDevices != nil {
		in, out := &in.HostDevices, &out.HostDevices
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LinstorStoragePoolSource.
func (in *LinstorStoragePoolSource) DeepCopy() *LinstorStoragePoolSource {
	if in == nil {
		return nil
	}
	out := new(LinstorStoragePoolSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Patch) DeepCopyInto(out *Patch) {
	*out = *in
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(Selector)
		**out = **in
	}
	if in.Options != nil {
		in, out := &in.Options, &out.Options
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Patch.
func (in *Patch) DeepCopy() *Patch {
	if in == nil {
		return nil
	}
	out := new(Patch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Selector) DeepCopyInto(out *Selector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Selector.
func (in *Selector) DeepCopy() *Selector {
	if in == nil {
		return nil
	}
	out := new(Selector)
	in.DeepCopyInto(out)
	return out
}