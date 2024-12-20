//go:build !ignore_autogenerated

/*
Copyright 2021.

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
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CSIDriverSpec) DeepCopyInto(out *CSIDriverSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CSIDriverSpec.
func (in *CSIDriverSpec) DeepCopy() *CSIDriverSpec {
	if in == nil {
		return nil
	}
	out := new(CSIDriverSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerStorageModule) DeepCopyInto(out *ContainerStorageModule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerStorageModule.
func (in *ContainerStorageModule) DeepCopy() *ContainerStorageModule {
	if in == nil {
		return nil
	}
	out := new(ContainerStorageModule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContainerStorageModule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerStorageModuleList) DeepCopyInto(out *ContainerStorageModuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ContainerStorageModule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerStorageModuleList.
func (in *ContainerStorageModuleList) DeepCopy() *ContainerStorageModuleList {
	if in == nil {
		return nil
	}
	out := new(ContainerStorageModuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContainerStorageModuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerStorageModuleSpec) DeepCopyInto(out *ContainerStorageModuleSpec) {
	*out = *in
	in.Driver.DeepCopyInto(&out.Driver)
	if in.Modules != nil {
		in, out := &in.Modules, &out.Modules
		*out = make([]Module, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerStorageModuleSpec.
func (in *ContainerStorageModuleSpec) DeepCopy() *ContainerStorageModuleSpec {
	if in == nil {
		return nil
	}
	out := new(ContainerStorageModuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerStorageModuleStatus) DeepCopyInto(out *ContainerStorageModuleStatus) {
	*out = *in
	out.ControllerStatus = in.ControllerStatus
	out.NodeStatus = in.NodeStatus
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerStorageModuleStatus.
func (in *ContainerStorageModuleStatus) DeepCopy() *ContainerStorageModuleStatus {
	if in == nil {
		return nil
	}
	out := new(ContainerStorageModuleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerTemplate) DeepCopyInto(out *ContainerTemplate) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Envs != nil {
		in, out := &in.Envs, &out.Envs
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ProxyServerIngress != nil {
		in, out := &in.ProxyServerIngress, &out.ProxyServerIngress
		*out = make([]ProxyServerIngress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Vaults != nil {
		in, out := &in.Vaults, &out.Vaults
		*out = make([]Vault, len(*in))
		copy(*out, *in)
	}
	if in.ComponentCred != nil {
		in, out := &in.ComponentCred, &out.ComponentCred
		*out = make([]Credential, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerTemplate.
func (in *ContainerTemplate) DeepCopy() *ContainerTemplate {
	if in == nil {
		return nil
	}
	out := new(ContainerTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Credential) DeepCopyInto(out *Credential) {
	*out = *in
	out.SecretContents = in.SecretContents
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Credential.
func (in *Credential) DeepCopy() *Credential {
	if in == nil {
		return nil
	}
	out := new(Credential)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Credkey) DeepCopyInto(out *Credkey) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Credkey.
func (in *Credkey) DeepCopy() *Credkey {
	if in == nil {
		return nil
	}
	out := new(Credkey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Driver) DeepCopyInto(out *Driver) {
	*out = *in
	if in.CSIDriverSpec != nil {
		in, out := &in.CSIDriverSpec, &out.CSIDriverSpec
		*out = new(CSIDriverSpec)
		**out = **in
	}
	if in.Common != nil {
		in, out := &in.Common, &out.Common
		*out = new(ContainerTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.Controller != nil {
		in, out := &in.Controller, &out.Controller
		*out = new(ContainerTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.Node != nil {
		in, out := &in.Node, &out.Node
		*out = new(ContainerTemplate)
		(*in).DeepCopyInto(*out)
	}
	if in.SideCars != nil {
		in, out := &in.SideCars, &out.SideCars
		*out = make([]ContainerTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitContainers != nil {
		in, out := &in.InitContainers, &out.InitContainers
		*out = make([]ContainerTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SnapshotClass != nil {
		in, out := &in.SnapshotClass, &out.SnapshotClass
		*out = make([]SnapshotClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ForceRemoveDriver != nil {
		in, out := &in.ForceRemoveDriver, &out.ForceRemoveDriver
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Driver.
func (in *Driver) DeepCopy() *Driver {
	if in == nil {
		return nil
	}
	out := new(Driver)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Module) DeepCopyInto(out *Module) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]ContainerTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.InitContainer != nil {
		in, out := &in.InitContainer, &out.InitContainer
		*out = make([]ContainerTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Module.
func (in *Module) DeepCopy() *Module {
	if in == nil {
		return nil
	}
	out := new(Module)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodStatus) DeepCopyInto(out *PodStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodStatus.
func (in *PodStatus) DeepCopy() *PodStatus {
	if in == nil {
		return nil
	}
	out := new(PodStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyServerIngress) DeepCopyInto(out *ProxyServerIngress) {
	*out = *in
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyServerIngress.
func (in *ProxyServerIngress) DeepCopy() *ProxyServerIngress {
	if in == nil {
		return nil
	}
	out := new(ProxyServerIngress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SnapshotClass) DeepCopyInto(out *SnapshotClass) {
	*out = *in
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SnapshotClass.
func (in *SnapshotClass) DeepCopy() *SnapshotClass {
	if in == nil {
		return nil
	}
	out := new(SnapshotClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vault) DeepCopyInto(out *Vault) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vault.
func (in *Vault) DeepCopy() *Vault {
	if in == nil {
		return nil
	}
	out := new(Vault)
	in.DeepCopyInto(out)
	return out
}
