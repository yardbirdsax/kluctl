//go:build !ignore_autogenerated

/*
Copyright 2023.

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

package result

import (
	"github.com/kluctl/kluctl/v2/pkg/types"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BaseObject) DeepCopyInto(out *BaseObject) {
	*out = *in
	out.Ref = in.Ref
	if in.Changes != nil {
		in, out := &in.Changes, &out.Changes
		*out = make([]Change, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BaseObject.
func (in *BaseObject) DeepCopy() *BaseObject {
	if in == nil {
		return nil
	}
	out := new(BaseObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Change) DeepCopyInto(out *Change) {
	*out = *in
	if in.OldValue != nil {
		in, out := &in.OldValue, &out.OldValue
		*out = new(v1.JSON)
		(*in).DeepCopyInto(*out)
	}
	if in.NewValue != nil {
		in, out := &in.NewValue, &out.NewValue
		*out = new(v1.JSON)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Change.
func (in *Change) DeepCopy() *Change {
	if in == nil {
		return nil
	}
	out := new(Change)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChangedObject) DeepCopyInto(out *ChangedObject) {
	*out = *in
	out.Ref = in.Ref
	if in.Changes != nil {
		in, out := &in.Changes, &out.Changes
		*out = make([]Change, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChangedObject.
func (in *ChangedObject) DeepCopy() *ChangedObject {
	if in == nil {
		return nil
	}
	out := new(ChangedObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterInfo) DeepCopyInto(out *ClusterInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterInfo.
func (in *ClusterInfo) DeepCopy() *ClusterInfo {
	if in == nil {
		return nil
	}
	out := new(ClusterInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommandInfo) DeepCopyInto(out *CommandInfo) {
	*out = *in
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.EndTime.DeepCopyInto(&out.EndTime)
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = (*in).DeepCopy()
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]types.FixedImage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.IncludeTags != nil {
		in, out := &in.IncludeTags, &out.IncludeTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludeTags != nil {
		in, out := &in.ExcludeTags, &out.ExcludeTags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.IncludeDeploymentDirs != nil {
		in, out := &in.IncludeDeploymentDirs, &out.IncludeDeploymentDirs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExcludeDeploymentDirs != nil {
		in, out := &in.ExcludeDeploymentDirs, &out.ExcludeDeploymentDirs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommandInfo.
func (in *CommandInfo) DeepCopy() *CommandInfo {
	if in == nil {
		return nil
	}
	out := new(CommandInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommandResult) DeepCopyInto(out *CommandResult) {
	*out = *in
	out.ProjectKey = in.ProjectKey
	out.TargetKey = in.TargetKey
	in.Target.DeepCopyInto(&out.Target)
	in.Command.DeepCopyInto(&out.Command)
	if in.KluctlDeployment != nil {
		in, out := &in.KluctlDeployment, &out.KluctlDeployment
		*out = new(KluctlDeploymentInfo)
		**out = **in
	}
	if in.OverridesPatch != nil {
		in, out := &in.OverridesPatch, &out.OverridesPatch
		*out = (*in).DeepCopy()
	}
	in.GitInfo.DeepCopyInto(&out.GitInfo)
	out.ClusterInfo = in.ClusterInfo
	if in.Deployment != nil {
		in, out := &in.Deployment, &out.Deployment
		*out = new(types.DeploymentProjectConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = make([]ResultObject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]DeploymentError, len(*in))
		copy(*out, *in)
	}
	if in.Warnings != nil {
		in, out := &in.Warnings, &out.Warnings
		*out = make([]DeploymentError, len(*in))
		copy(*out, *in)
	}
	if in.SeenImages != nil {
		in, out := &in.SeenImages, &out.SeenImages
		*out = make([]types.FixedImage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommandResult.
func (in *CommandResult) DeepCopy() *CommandResult {
	if in == nil {
		return nil
	}
	out := new(CommandResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommandResultSummary) DeepCopyInto(out *CommandResultSummary) {
	*out = *in
	out.ProjectKey = in.ProjectKey
	out.TargetKey = in.TargetKey
	in.Target.DeepCopyInto(&out.Target)
	in.Command.DeepCopyInto(&out.Command)
	if in.KluctlDeployment != nil {
		in, out := &in.KluctlDeployment, &out.KluctlDeployment
		*out = new(KluctlDeploymentInfo)
		**out = **in
	}
	in.GitInfo.DeepCopyInto(&out.GitInfo)
	out.ClusterInfo = in.ClusterInfo
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]DeploymentError, len(*in))
		copy(*out, *in)
	}
	if in.Warnings != nil {
		in, out := &in.Warnings, &out.Warnings
		*out = make([]DeploymentError, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommandResultSummary.
func (in *CommandResultSummary) DeepCopy() *CommandResultSummary {
	if in == nil {
		return nil
	}
	out := new(CommandResultSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompactedCommandResult) DeepCopyInto(out *CompactedCommandResult) {
	*out = *in
	in.CommandResult.DeepCopyInto(&out.CommandResult)
	if in.CompactedObjects != nil {
		in, out := &in.CompactedObjects, &out.CompactedObjects
		*out = make(CompactedObjects, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompactedCommandResult.
func (in *CompactedCommandResult) DeepCopy() *CompactedCommandResult {
	if in == nil {
		return nil
	}
	out := new(CompactedCommandResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CompactedObject) DeepCopyInto(out *CompactedObject) {
	*out = *in
	in.BaseObject.DeepCopyInto(&out.BaseObject)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompactedObject.
func (in *CompactedObject) DeepCopy() *CompactedObject {
	if in == nil {
		return nil
	}
	out := new(CompactedObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in CompactedObjects) DeepCopyInto(out *CompactedObjects) {
	{
		in := &in
		*out = make(CompactedObjects, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CompactedObjects.
func (in CompactedObjects) DeepCopy() CompactedObjects {
	if in == nil {
		return nil
	}
	out := new(CompactedObjects)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentError) DeepCopyInto(out *DeploymentError) {
	*out = *in
	out.Ref = in.Ref
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentError.
func (in *DeploymentError) DeepCopy() *DeploymentError {
	if in == nil {
		return nil
	}
	out := new(DeploymentError)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriftDetectionResult) DeepCopyInto(out *DriftDetectionResult) {
	*out = *in
	out.ProjectKey = in.ProjectKey
	out.TargetKey = in.TargetKey
	if in.KluctlDeployment != nil {
		in, out := &in.KluctlDeployment, &out.KluctlDeployment
		*out = new(KluctlDeploymentInfo)
		**out = **in
	}
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.EndTime.DeepCopyInto(&out.EndTime)
	if in.Warnings != nil {
		in, out := &in.Warnings, &out.Warnings
		*out = make([]DeploymentError, len(*in))
		copy(*out, *in)
	}
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]DeploymentError, len(*in))
		copy(*out, *in)
	}
	if in.Objects != nil {
		in, out := &in.Objects, &out.Objects
		*out = make([]DriftedObject, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriftDetectionResult.
func (in *DriftDetectionResult) DeepCopy() *DriftDetectionResult {
	if in == nil {
		return nil
	}
	out := new(DriftDetectionResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DriftedObject) DeepCopyInto(out *DriftedObject) {
	*out = *in
	in.BaseObject.DeepCopyInto(&out.BaseObject)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DriftedObject.
func (in *DriftedObject) DeepCopy() *DriftedObject {
	if in == nil {
		return nil
	}
	out := new(DriftedObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitInfo) DeepCopyInto(out *GitInfo) {
	*out = *in
	if in.Url != nil {
		in, out := &in.Url, &out.Url
		*out = (*in).DeepCopy()
	}
	if in.Ref != nil {
		in, out := &in.Ref, &out.Ref
		*out = new(types.GitRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitInfo.
func (in *GitInfo) DeepCopy() *GitInfo {
	if in == nil {
		return nil
	}
	out := new(GitInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KluctlDeploymentInfo) DeepCopyInto(out *KluctlDeploymentInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KluctlDeploymentInfo.
func (in *KluctlDeploymentInfo) DeepCopy() *KluctlDeploymentInfo {
	if in == nil {
		return nil
	}
	out := new(KluctlDeploymentInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectKey) DeepCopyInto(out *ProjectKey) {
	*out = *in
	out.GitRepoKey = in.GitRepoKey
	out.OciRepoKey = in.OciRepoKey
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectKey.
func (in *ProjectKey) DeepCopy() *ProjectKey {
	if in == nil {
		return nil
	}
	out := new(ProjectKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResultObject) DeepCopyInto(out *ResultObject) {
	*out = *in
	in.BaseObject.DeepCopyInto(&out.BaseObject)
	if in.Rendered != nil {
		in, out := &in.Rendered, &out.Rendered
		*out = (*in).DeepCopy()
	}
	if in.Remote != nil {
		in, out := &in.Remote, &out.Remote
		*out = (*in).DeepCopy()
	}
	if in.Applied != nil {
		in, out := &in.Applied, &out.Applied
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResultObject.
func (in *ResultObject) DeepCopy() *ResultObject {
	if in == nil {
		return nil
	}
	out := new(ResultObject)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TargetKey) DeepCopyInto(out *TargetKey) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TargetKey.
func (in *TargetKey) DeepCopy() *TargetKey {
	if in == nil {
		return nil
	}
	out := new(TargetKey)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidateResult) DeepCopyInto(out *ValidateResult) {
	*out = *in
	out.ProjectKey = in.ProjectKey
	out.TargetKey = in.TargetKey
	if in.KluctlDeployment != nil {
		in, out := &in.KluctlDeployment, &out.KluctlDeployment
		*out = new(KluctlDeploymentInfo)
		**out = **in
	}
	if in.OverridesPatch != nil {
		in, out := &in.OverridesPatch, &out.OverridesPatch
		*out = (*in).DeepCopy()
	}
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.EndTime.DeepCopyInto(&out.EndTime)
	if in.Warnings != nil {
		in, out := &in.Warnings, &out.Warnings
		*out = make([]DeploymentError, len(*in))
		copy(*out, *in)
	}
	if in.Errors != nil {
		in, out := &in.Errors, &out.Errors
		*out = make([]DeploymentError, len(*in))
		copy(*out, *in)
	}
	if in.Results != nil {
		in, out := &in.Results, &out.Results
		*out = make([]ValidateResultEntry, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidateResult.
func (in *ValidateResult) DeepCopy() *ValidateResult {
	if in == nil {
		return nil
	}
	out := new(ValidateResult)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidateResultEntry) DeepCopyInto(out *ValidateResultEntry) {
	*out = *in
	out.Ref = in.Ref
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidateResultEntry.
func (in *ValidateResultEntry) DeepCopy() *ValidateResultEntry {
	if in == nil {
		return nil
	}
	out := new(ValidateResultEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValidateResultSummary) DeepCopyInto(out *ValidateResultSummary) {
	*out = *in
	out.ProjectKey = in.ProjectKey
	out.TargetKey = in.TargetKey
	if in.KluctlDeployment != nil {
		in, out := &in.KluctlDeployment, &out.KluctlDeployment
		*out = new(KluctlDeploymentInfo)
		**out = **in
	}
	in.StartTime.DeepCopyInto(&out.StartTime)
	in.EndTime.DeepCopyInto(&out.EndTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValidateResultSummary.
func (in *ValidateResultSummary) DeepCopy() *ValidateResultSummary {
	if in == nil {
		return nil
	}
	out := new(ValidateResultSummary)
	in.DeepCopyInto(out)
	return out
}
