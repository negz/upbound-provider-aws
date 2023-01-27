//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Upbound Inc.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACLConfigurationObservation) DeepCopyInto(out *ACLConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACLConfigurationObservation.
func (in *ACLConfigurationObservation) DeepCopy() *ACLConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(ACLConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ACLConfigurationParameters) DeepCopyInto(out *ACLConfigurationParameters) {
	*out = *in
	if in.S3ACLOption != nil {
		in, out := &in.S3ACLOption, &out.S3ACLOption
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ACLConfigurationParameters.
func (in *ACLConfigurationParameters) DeepCopy() *ACLConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(ACLConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationObservation) DeepCopyInto(out *ConfigurationObservation) {
	*out = *in
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = make([]EngineVersionObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationObservation.
func (in *ConfigurationObservation) DeepCopy() *ConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(ConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigurationParameters) DeepCopyInto(out *ConfigurationParameters) {
	*out = *in
	if in.BytesScannedCutoffPerQuery != nil {
		in, out := &in.BytesScannedCutoffPerQuery, &out.BytesScannedCutoffPerQuery
		*out = new(float64)
		**out = **in
	}
	if in.EnforceWorkgroupConfiguration != nil {
		in, out := &in.EnforceWorkgroupConfiguration, &out.EnforceWorkgroupConfiguration
		*out = new(bool)
		**out = **in
	}
	if in.EngineVersion != nil {
		in, out := &in.EngineVersion, &out.EngineVersion
		*out = make([]EngineVersionParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExecutionRole != nil {
		in, out := &in.ExecutionRole, &out.ExecutionRole
		*out = new(string)
		**out = **in
	}
	if in.PublishCloudwatchMetricsEnabled != nil {
		in, out := &in.PublishCloudwatchMetricsEnabled, &out.PublishCloudwatchMetricsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RequesterPaysEnabled != nil {
		in, out := &in.RequesterPaysEnabled, &out.RequesterPaysEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ResultConfiguration != nil {
		in, out := &in.ResultConfiguration, &out.ResultConfiguration
		*out = make([]ResultConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigurationParameters.
func (in *ConfigurationParameters) DeepCopy() *ConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(ConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCatalog) DeepCopyInto(out *DataCatalog) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCatalog.
func (in *DataCatalog) DeepCopy() *DataCatalog {
	if in == nil {
		return nil
	}
	out := new(DataCatalog)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataCatalog) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCatalogList) DeepCopyInto(out *DataCatalogList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DataCatalog, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCatalogList.
func (in *DataCatalogList) DeepCopy() *DataCatalogList {
	if in == nil {
		return nil
	}
	out := new(DataCatalogList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DataCatalogList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCatalogObservation) DeepCopyInto(out *DataCatalogObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCatalogObservation.
func (in *DataCatalogObservation) DeepCopy() *DataCatalogObservation {
	if in == nil {
		return nil
	}
	out := new(DataCatalogObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCatalogParameters) DeepCopyInto(out *DataCatalogParameters) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Parameters != nil {
		in, out := &in.Parameters, &out.Parameters
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCatalogParameters.
func (in *DataCatalogParameters) DeepCopy() *DataCatalogParameters {
	if in == nil {
		return nil
	}
	out := new(DataCatalogParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCatalogSpec) DeepCopyInto(out *DataCatalogSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCatalogSpec.
func (in *DataCatalogSpec) DeepCopy() *DataCatalogSpec {
	if in == nil {
		return nil
	}
	out := new(DataCatalogSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataCatalogStatus) DeepCopyInto(out *DataCatalogStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataCatalogStatus.
func (in *DataCatalogStatus) DeepCopy() *DataCatalogStatus {
	if in == nil {
		return nil
	}
	out := new(DataCatalogStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Database) DeepCopyInto(out *Database) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Database.
func (in *Database) DeepCopy() *Database {
	if in == nil {
		return nil
	}
	out := new(Database)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Database) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseList) DeepCopyInto(out *DatabaseList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Database, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseList.
func (in *DatabaseList) DeepCopy() *DatabaseList {
	if in == nil {
		return nil
	}
	out := new(DatabaseList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseObservation) DeepCopyInto(out *DatabaseObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseObservation.
func (in *DatabaseObservation) DeepCopy() *DatabaseObservation {
	if in == nil {
		return nil
	}
	out := new(DatabaseObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseParameters) DeepCopyInto(out *DatabaseParameters) {
	*out = *in
	if in.ACLConfiguration != nil {
		in, out := &in.ACLConfiguration, &out.ACLConfiguration
		*out = make([]ACLConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Bucket != nil {
		in, out := &in.Bucket, &out.Bucket
		*out = new(string)
		**out = **in
	}
	if in.BucketRef != nil {
		in, out := &in.BucketRef, &out.BucketRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.BucketSelector != nil {
		in, out := &in.BucketSelector, &out.BucketSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Comment != nil {
		in, out := &in.Comment, &out.Comment
		*out = new(string)
		**out = **in
	}
	if in.EncryptionConfiguration != nil {
		in, out := &in.EncryptionConfiguration, &out.EncryptionConfiguration
		*out = make([]EncryptionConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExpectedBucketOwner != nil {
		in, out := &in.ExpectedBucketOwner, &out.ExpectedBucketOwner
		*out = new(string)
		**out = **in
	}
	if in.ForceDestroy != nil {
		in, out := &in.ForceDestroy, &out.ForceDestroy
		*out = new(bool)
		**out = **in
	}
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseParameters.
func (in *DatabaseParameters) DeepCopy() *DatabaseParameters {
	if in == nil {
		return nil
	}
	out := new(DatabaseParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseSpec) DeepCopyInto(out *DatabaseSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseSpec.
func (in *DatabaseSpec) DeepCopy() *DatabaseSpec {
	if in == nil {
		return nil
	}
	out := new(DatabaseSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseStatus) DeepCopyInto(out *DatabaseStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseStatus.
func (in *DatabaseStatus) DeepCopy() *DatabaseStatus {
	if in == nil {
		return nil
	}
	out := new(DatabaseStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionConfigurationObservation) DeepCopyInto(out *EncryptionConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionConfigurationObservation.
func (in *EncryptionConfigurationObservation) DeepCopy() *EncryptionConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(EncryptionConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionConfigurationParameters) DeepCopyInto(out *EncryptionConfigurationParameters) {
	*out = *in
	if in.EncryptionOption != nil {
		in, out := &in.EncryptionOption, &out.EncryptionOption
		*out = new(string)
		**out = **in
	}
	if in.KMSKey != nil {
		in, out := &in.KMSKey, &out.KMSKey
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionConfigurationParameters.
func (in *EncryptionConfigurationParameters) DeepCopy() *EncryptionConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(EncryptionConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EngineVersionObservation) DeepCopyInto(out *EngineVersionObservation) {
	*out = *in
	if in.EffectiveEngineVersion != nil {
		in, out := &in.EffectiveEngineVersion, &out.EffectiveEngineVersion
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EngineVersionObservation.
func (in *EngineVersionObservation) DeepCopy() *EngineVersionObservation {
	if in == nil {
		return nil
	}
	out := new(EngineVersionObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EngineVersionParameters) DeepCopyInto(out *EngineVersionParameters) {
	*out = *in
	if in.SelectedEngineVersion != nil {
		in, out := &in.SelectedEngineVersion, &out.SelectedEngineVersion
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EngineVersionParameters.
func (in *EngineVersionParameters) DeepCopy() *EngineVersionParameters {
	if in == nil {
		return nil
	}
	out := new(EngineVersionParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedQuery) DeepCopyInto(out *NamedQuery) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedQuery.
func (in *NamedQuery) DeepCopy() *NamedQuery {
	if in == nil {
		return nil
	}
	out := new(NamedQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamedQuery) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedQueryList) DeepCopyInto(out *NamedQueryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NamedQuery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedQueryList.
func (in *NamedQueryList) DeepCopy() *NamedQueryList {
	if in == nil {
		return nil
	}
	out := new(NamedQueryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NamedQueryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedQueryObservation) DeepCopyInto(out *NamedQueryObservation) {
	*out = *in
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedQueryObservation.
func (in *NamedQueryObservation) DeepCopy() *NamedQueryObservation {
	if in == nil {
		return nil
	}
	out := new(NamedQueryObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedQueryParameters) DeepCopyInto(out *NamedQueryParameters) {
	*out = *in
	if in.Database != nil {
		in, out := &in.Database, &out.Database
		*out = new(string)
		**out = **in
	}
	if in.DatabaseRef != nil {
		in, out := &in.DatabaseRef, &out.DatabaseRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.DatabaseSelector != nil {
		in, out := &in.DatabaseSelector, &out.DatabaseSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Workgroup != nil {
		in, out := &in.Workgroup, &out.Workgroup
		*out = new(string)
		**out = **in
	}
	if in.WorkgroupRef != nil {
		in, out := &in.WorkgroupRef, &out.WorkgroupRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkgroupSelector != nil {
		in, out := &in.WorkgroupSelector, &out.WorkgroupSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedQueryParameters.
func (in *NamedQueryParameters) DeepCopy() *NamedQueryParameters {
	if in == nil {
		return nil
	}
	out := new(NamedQueryParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedQuerySpec) DeepCopyInto(out *NamedQuerySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedQuerySpec.
func (in *NamedQuerySpec) DeepCopy() *NamedQuerySpec {
	if in == nil {
		return nil
	}
	out := new(NamedQuerySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamedQueryStatus) DeepCopyInto(out *NamedQueryStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamedQueryStatus.
func (in *NamedQueryStatus) DeepCopy() *NamedQueryStatus {
	if in == nil {
		return nil
	}
	out := new(NamedQueryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResultConfigurationACLConfigurationObservation) DeepCopyInto(out *ResultConfigurationACLConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResultConfigurationACLConfigurationObservation.
func (in *ResultConfigurationACLConfigurationObservation) DeepCopy() *ResultConfigurationACLConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(ResultConfigurationACLConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResultConfigurationACLConfigurationParameters) DeepCopyInto(out *ResultConfigurationACLConfigurationParameters) {
	*out = *in
	if in.S3ACLOption != nil {
		in, out := &in.S3ACLOption, &out.S3ACLOption
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResultConfigurationACLConfigurationParameters.
func (in *ResultConfigurationACLConfigurationParameters) DeepCopy() *ResultConfigurationACLConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(ResultConfigurationACLConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResultConfigurationEncryptionConfigurationObservation) DeepCopyInto(out *ResultConfigurationEncryptionConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResultConfigurationEncryptionConfigurationObservation.
func (in *ResultConfigurationEncryptionConfigurationObservation) DeepCopy() *ResultConfigurationEncryptionConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(ResultConfigurationEncryptionConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResultConfigurationEncryptionConfigurationParameters) DeepCopyInto(out *ResultConfigurationEncryptionConfigurationParameters) {
	*out = *in
	if in.EncryptionOption != nil {
		in, out := &in.EncryptionOption, &out.EncryptionOption
		*out = new(string)
		**out = **in
	}
	if in.KMSKeyArn != nil {
		in, out := &in.KMSKeyArn, &out.KMSKeyArn
		*out = new(string)
		**out = **in
	}
	if in.KMSKeyArnRef != nil {
		in, out := &in.KMSKeyArnRef, &out.KMSKeyArnRef
		*out = new(v1.Reference)
		(*in).DeepCopyInto(*out)
	}
	if in.KMSKeyArnSelector != nil {
		in, out := &in.KMSKeyArnSelector, &out.KMSKeyArnSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResultConfigurationEncryptionConfigurationParameters.
func (in *ResultConfigurationEncryptionConfigurationParameters) DeepCopy() *ResultConfigurationEncryptionConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(ResultConfigurationEncryptionConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResultConfigurationObservation) DeepCopyInto(out *ResultConfigurationObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResultConfigurationObservation.
func (in *ResultConfigurationObservation) DeepCopy() *ResultConfigurationObservation {
	if in == nil {
		return nil
	}
	out := new(ResultConfigurationObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResultConfigurationParameters) DeepCopyInto(out *ResultConfigurationParameters) {
	*out = *in
	if in.ACLConfiguration != nil {
		in, out := &in.ACLConfiguration, &out.ACLConfiguration
		*out = make([]ResultConfigurationACLConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.EncryptionConfiguration != nil {
		in, out := &in.EncryptionConfiguration, &out.EncryptionConfiguration
		*out = make([]ResultConfigurationEncryptionConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExpectedBucketOwner != nil {
		in, out := &in.ExpectedBucketOwner, &out.ExpectedBucketOwner
		*out = new(string)
		**out = **in
	}
	if in.OutputLocation != nil {
		in, out := &in.OutputLocation, &out.OutputLocation
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResultConfigurationParameters.
func (in *ResultConfigurationParameters) DeepCopy() *ResultConfigurationParameters {
	if in == nil {
		return nil
	}
	out := new(ResultConfigurationParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workgroup) DeepCopyInto(out *Workgroup) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workgroup.
func (in *Workgroup) DeepCopy() *Workgroup {
	if in == nil {
		return nil
	}
	out := new(Workgroup)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Workgroup) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkgroupList) DeepCopyInto(out *WorkgroupList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Workgroup, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkgroupList.
func (in *WorkgroupList) DeepCopy() *WorkgroupList {
	if in == nil {
		return nil
	}
	out := new(WorkgroupList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkgroupList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkgroupObservation) DeepCopyInto(out *WorkgroupObservation) {
	*out = *in
	if in.Arn != nil {
		in, out := &in.Arn, &out.Arn
		*out = new(string)
		**out = **in
	}
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = make([]ConfigurationObservation, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.TagsAll != nil {
		in, out := &in.TagsAll, &out.TagsAll
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkgroupObservation.
func (in *WorkgroupObservation) DeepCopy() *WorkgroupObservation {
	if in == nil {
		return nil
	}
	out := new(WorkgroupObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkgroupParameters) DeepCopyInto(out *WorkgroupParameters) {
	*out = *in
	if in.Configuration != nil {
		in, out := &in.Configuration, &out.Configuration
		*out = make([]ConfigurationParameters, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.ForceDestroy != nil {
		in, out := &in.ForceDestroy, &out.ForceDestroy
		*out = new(bool)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make(map[string]*string, len(*in))
		for key, val := range *in {
			var outVal *string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(string)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkgroupParameters.
func (in *WorkgroupParameters) DeepCopy() *WorkgroupParameters {
	if in == nil {
		return nil
	}
	out := new(WorkgroupParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkgroupSpec) DeepCopyInto(out *WorkgroupSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkgroupSpec.
func (in *WorkgroupSpec) DeepCopy() *WorkgroupSpec {
	if in == nil {
		return nil
	}
	out := new(WorkgroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkgroupStatus) DeepCopyInto(out *WorkgroupStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkgroupStatus.
func (in *WorkgroupStatus) DeepCopy() *WorkgroupStatus {
	if in == nil {
		return nil
	}
	out := new(WorkgroupStatus)
	in.DeepCopyInto(out)
	return out
}
