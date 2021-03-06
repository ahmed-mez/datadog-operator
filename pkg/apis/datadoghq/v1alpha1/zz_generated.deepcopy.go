// +build !ignore_autogenerated

// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-2019 Datadog, Inc.

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	datadoghqv1alpha1 "github.com/datadog/extendeddaemonset/pkg/apis/datadoghq/v1alpha1"
	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	intstr "k8s.io/apimachinery/pkg/util/intstr"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APMSpec) DeepCopyInto(out *APMSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.HostPort != nil {
		in, out := &in.HostPort, &out.HostPort
		*out = new(int32)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APMSpec.
func (in *APMSpec) DeepCopy() *APMSpec {
	if in == nil {
		return nil
	}
	out := new(APMSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AgentCredentials) DeepCopyInto(out *AgentCredentials) {
	*out = *in
	if in.UseSecretBackend != nil {
		in, out := &in.UseSecretBackend, &out.UseSecretBackend
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentCredentials.
func (in *AgentCredentials) DeepCopy() *AgentCredentials {
	if in == nil {
		return nil
	}
	out := new(AgentCredentials)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CRISocketConfig) DeepCopyInto(out *CRISocketConfig) {
	*out = *in
	if in.UseCriSocketVolume != nil {
		in, out := &in.UseCriSocketVolume, &out.UseCriSocketVolume
		*out = new(bool)
		**out = **in
	}
	if in.CriSocketPath != nil {
		in, out := &in.CriSocketPath, &out.CriSocketPath
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CRISocketConfig.
func (in *CRISocketConfig) DeepCopy() *CRISocketConfig {
	if in == nil {
		return nil
	}
	out := new(CRISocketConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterAgentConfig) DeepCopyInto(out *ClusterAgentConfig) {
	*out = *in
	if in.MetricsProviderEnabled != nil {
		in, out := &in.MetricsProviderEnabled, &out.MetricsProviderEnabled
		*out = new(bool)
		**out = **in
	}
	if in.MetricsProviderPort != nil {
		in, out := &in.MetricsProviderPort, &out.MetricsProviderPort
		*out = new(int32)
		**out = **in
	}
	if in.ClusterChecksRunnerEnabled != nil {
		in, out := &in.ClusterChecksRunnerEnabled, &out.ClusterChecksRunnerEnabled
		*out = new(bool)
		**out = **in
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.Confd != nil {
		in, out := &in.Confd, &out.Confd
		*out = new(ConfigDirSpec)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterAgentConfig.
func (in *ClusterAgentConfig) DeepCopy() *ClusterAgentConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterAgentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterChecksRunnerConfig) DeepCopyInto(out *ClusterChecksRunnerConfig) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterChecksRunnerConfig.
func (in *ClusterChecksRunnerConfig) DeepCopy() *ClusterChecksRunnerConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterChecksRunnerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigDirSpec) DeepCopyInto(out *ConfigDirSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigDirSpec.
func (in *ConfigDirSpec) DeepCopy() *ConfigDirSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigDirSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigFileConfigMapSpec) DeepCopyInto(out *ConfigFileConfigMapSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigFileConfigMapSpec.
func (in *ConfigFileConfigMapSpec) DeepCopy() *ConfigFileConfigMapSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigFileConfigMapSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomConfigSpec) DeepCopyInto(out *CustomConfigSpec) {
	*out = *in
	if in.ConfigData != nil {
		in, out := &in.ConfigData, &out.ConfigData
		*out = new(string)
		**out = **in
	}
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(ConfigFileConfigMapSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomConfigSpec.
func (in *CustomConfigSpec) DeepCopy() *CustomConfigSpec {
	if in == nil {
		return nil
	}
	out := new(CustomConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonSetDeploymentStrategy) DeepCopyInto(out *DaemonSetDeploymentStrategy) {
	*out = *in
	if in.UpdateStrategyType != nil {
		in, out := &in.UpdateStrategyType, &out.UpdateStrategyType
		*out = new(appsv1.DaemonSetUpdateStrategyType)
		**out = **in
	}
	in.RollingUpdate.DeepCopyInto(&out.RollingUpdate)
	if in.Canary != nil {
		in, out := &in.Canary, &out.Canary
		*out = new(datadoghqv1alpha1.ExtendedDaemonSetSpecStrategyCanary)
		(*in).DeepCopyInto(*out)
	}
	if in.ReconcileFrequency != nil {
		in, out := &in.ReconcileFrequency, &out.ReconcileFrequency
		*out = new(metav1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonSetDeploymentStrategy.
func (in *DaemonSetDeploymentStrategy) DeepCopy() *DaemonSetDeploymentStrategy {
	if in == nil {
		return nil
	}
	out := new(DaemonSetDeploymentStrategy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonSetRollingUpdateSpec) DeepCopyInto(out *DaemonSetRollingUpdateSpec) {
	*out = *in
	if in.MaxUnavailable != nil {
		in, out := &in.MaxUnavailable, &out.MaxUnavailable
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.MaxPodSchedulerFailure != nil {
		in, out := &in.MaxPodSchedulerFailure, &out.MaxPodSchedulerFailure
		*out = new(intstr.IntOrString)
		**out = **in
	}
	if in.MaxParallelPodCreation != nil {
		in, out := &in.MaxParallelPodCreation, &out.MaxParallelPodCreation
		*out = new(int32)
		**out = **in
	}
	if in.SlowStartIntervalDuration != nil {
		in, out := &in.SlowStartIntervalDuration, &out.SlowStartIntervalDuration
		*out = new(metav1.Duration)
		**out = **in
	}
	if in.SlowStartAdditiveIncrease != nil {
		in, out := &in.SlowStartAdditiveIncrease, &out.SlowStartAdditiveIncrease
		*out = new(intstr.IntOrString)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonSetRollingUpdateSpec.
func (in *DaemonSetRollingUpdateSpec) DeepCopy() *DaemonSetRollingUpdateSpec {
	if in == nil {
		return nil
	}
	out := new(DaemonSetRollingUpdateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DaemonSetStatus) DeepCopyInto(out *DaemonSetStatus) {
	*out = *in
	if in.LastUpdate != nil {
		in, out := &in.LastUpdate, &out.LastUpdate
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DaemonSetStatus.
func (in *DaemonSetStatus) DeepCopy() *DaemonSetStatus {
	if in == nil {
		return nil
	}
	out := new(DaemonSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogAgent) DeepCopyInto(out *DatadogAgent) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogAgent.
func (in *DatadogAgent) DeepCopy() *DatadogAgent {
	if in == nil {
		return nil
	}
	out := new(DatadogAgent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatadogAgent) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogAgentCondition) DeepCopyInto(out *DatadogAgentCondition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogAgentCondition.
func (in *DatadogAgentCondition) DeepCopy() *DatadogAgentCondition {
	if in == nil {
		return nil
	}
	out := new(DatadogAgentCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogAgentList) DeepCopyInto(out *DatadogAgentList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatadogAgent, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogAgentList.
func (in *DatadogAgentList) DeepCopy() *DatadogAgentList {
	if in == nil {
		return nil
	}
	out := new(DatadogAgentList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatadogAgentList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogAgentSpec) DeepCopyInto(out *DatadogAgentSpec) {
	*out = *in
	in.Credentials.DeepCopyInto(&out.Credentials)
	if in.Agent != nil {
		in, out := &in.Agent, &out.Agent
		*out = new(DatadogAgentSpecAgentSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterAgent != nil {
		in, out := &in.ClusterAgent, &out.ClusterAgent
		*out = new(DatadogAgentSpecClusterAgentSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterChecksRunner != nil {
		in, out := &in.ClusterChecksRunner, &out.ClusterChecksRunner
		*out = new(DatadogAgentSpecClusterChecksRunnerSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogAgentSpec.
func (in *DatadogAgentSpec) DeepCopy() *DatadogAgentSpec {
	if in == nil {
		return nil
	}
	out := new(DatadogAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogAgentSpecAgentSpec) DeepCopyInto(out *DatadogAgentSpecAgentSpec) {
	*out = *in
	if in.UseExtendedDaemonset != nil {
		in, out := &in.UseExtendedDaemonset, &out.UseExtendedDaemonset
		*out = new(bool)
		**out = **in
	}
	in.Image.DeepCopyInto(&out.Image)
	in.Config.DeepCopyInto(&out.Config)
	in.Rbac.DeepCopyInto(&out.Rbac)
	if in.DeploymentStrategy != nil {
		in, out := &in.DeploymentStrategy, &out.DeploymentStrategy
		*out = new(DaemonSetDeploymentStrategy)
		(*in).DeepCopyInto(*out)
	}
	if in.AdditionalAnnotations != nil {
		in, out := &in.AdditionalAnnotations, &out.AdditionalAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AdditionalLabels != nil {
		in, out := &in.AdditionalLabels, &out.AdditionalLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.DNSConfig != nil {
		in, out := &in.DNSConfig, &out.DNSConfig
		*out = new(v1.PodDNSConfig)
		(*in).DeepCopyInto(*out)
	}
	in.Apm.DeepCopyInto(&out.Apm)
	in.Log.DeepCopyInto(&out.Log)
	in.Process.DeepCopyInto(&out.Process)
	in.SystemProbe.DeepCopyInto(&out.SystemProbe)
	if in.CustomConfig != nil {
		in, out := &in.CustomConfig, &out.CustomConfig
		*out = new(CustomConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogAgentSpecAgentSpec.
func (in *DatadogAgentSpecAgentSpec) DeepCopy() *DatadogAgentSpecAgentSpec {
	if in == nil {
		return nil
	}
	out := new(DatadogAgentSpecAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogAgentSpecClusterAgentSpec) DeepCopyInto(out *DatadogAgentSpecClusterAgentSpec) {
	*out = *in
	in.Image.DeepCopyInto(&out.Image)
	in.Config.DeepCopyInto(&out.Config)
	if in.CustomConfig != nil {
		in, out := &in.CustomConfig, &out.CustomConfig
		*out = new(CustomConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Rbac.DeepCopyInto(&out.Rbac)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.AdditionalAnnotations != nil {
		in, out := &in.AdditionalAnnotations, &out.AdditionalAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AdditionalLabels != nil {
		in, out := &in.AdditionalLabels, &out.AdditionalLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
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
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogAgentSpecClusterAgentSpec.
func (in *DatadogAgentSpecClusterAgentSpec) DeepCopy() *DatadogAgentSpecClusterAgentSpec {
	if in == nil {
		return nil
	}
	out := new(DatadogAgentSpecClusterAgentSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogAgentSpecClusterChecksRunnerSpec) DeepCopyInto(out *DatadogAgentSpecClusterChecksRunnerSpec) {
	*out = *in
	in.Image.DeepCopyInto(&out.Image)
	in.Config.DeepCopyInto(&out.Config)
	if in.CustomConfig != nil {
		in, out := &in.CustomConfig, &out.CustomConfig
		*out = new(CustomConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	in.Rbac.DeepCopyInto(&out.Rbac)
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.AdditionalAnnotations != nil {
		in, out := &in.AdditionalAnnotations, &out.AdditionalAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.AdditionalLabels != nil {
		in, out := &in.AdditionalLabels, &out.AdditionalLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
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
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogAgentSpecClusterChecksRunnerSpec.
func (in *DatadogAgentSpecClusterChecksRunnerSpec) DeepCopy() *DatadogAgentSpecClusterChecksRunnerSpec {
	if in == nil {
		return nil
	}
	out := new(DatadogAgentSpecClusterChecksRunnerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatadogAgentStatus) DeepCopyInto(out *DatadogAgentStatus) {
	*out = *in
	if in.Agent != nil {
		in, out := &in.Agent, &out.Agent
		*out = new(DaemonSetStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterAgent != nil {
		in, out := &in.ClusterAgent, &out.ClusterAgent
		*out = new(DeploymentStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.ClusterChecksRunner != nil {
		in, out := &in.ClusterChecksRunner, &out.ClusterChecksRunner
		*out = new(DeploymentStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]DatadogAgentCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatadogAgentStatus.
func (in *DatadogAgentStatus) DeepCopy() *DatadogAgentStatus {
	if in == nil {
		return nil
	}
	out := new(DatadogAgentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeploymentStatus) DeepCopyInto(out *DeploymentStatus) {
	*out = *in
	if in.LastUpdate != nil {
		in, out := &in.LastUpdate, &out.LastUpdate
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeploymentStatus.
func (in *DeploymentStatus) DeepCopy() *DeploymentStatus {
	if in == nil {
		return nil
	}
	out := new(DeploymentStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DogstatsdConfig) DeepCopyInto(out *DogstatsdConfig) {
	*out = *in
	if in.DogstatsdOriginDetection != nil {
		in, out := &in.DogstatsdOriginDetection, &out.DogstatsdOriginDetection
		*out = new(bool)
		**out = **in
	}
	if in.UseDogStatsDSocketVolume != nil {
		in, out := &in.UseDogStatsDSocketVolume, &out.UseDogStatsDSocketVolume
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DogstatsdConfig.
func (in *DogstatsdConfig) DeepCopy() *DogstatsdConfig {
	if in == nil {
		return nil
	}
	out := new(DogstatsdConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageConfig) DeepCopyInto(out *ImageConfig) {
	*out = *in
	if in.PullPolicy != nil {
		in, out := &in.PullPolicy, &out.PullPolicy
		*out = new(v1.PullPolicy)
		**out = **in
	}
	if in.PullSecrets != nil {
		in, out := &in.PullSecrets, &out.PullSecrets
		*out = new([]v1.LocalObjectReference)
		if **in != nil {
			in, out := *in, *out
			*out = make([]v1.LocalObjectReference, len(*in))
			copy(*out, *in)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageConfig.
func (in *ImageConfig) DeepCopy() *ImageConfig {
	if in == nil {
		return nil
	}
	out := new(ImageConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogSpec) DeepCopyInto(out *LogSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.LogsConfigContainerCollectAll != nil {
		in, out := &in.LogsConfigContainerCollectAll, &out.LogsConfigContainerCollectAll
		*out = new(bool)
		**out = **in
	}
	if in.ContainerLogsPath != nil {
		in, out := &in.ContainerLogsPath, &out.ContainerLogsPath
		*out = new(string)
		**out = **in
	}
	if in.PodLogsPath != nil {
		in, out := &in.PodLogsPath, &out.PodLogsPath
		*out = new(string)
		**out = **in
	}
	if in.TempStoragePath != nil {
		in, out := &in.TempStoragePath, &out.TempStoragePath
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogSpec.
func (in *LogSpec) DeepCopy() *LogSpec {
	if in == nil {
		return nil
	}
	out := new(LogSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeAgentConfig) DeepCopyInto(out *NodeAgentConfig) {
	*out = *in
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.DDUrl != nil {
		in, out := &in.DDUrl, &out.DDUrl
		*out = new(string)
		**out = **in
	}
	if in.LogLevel != nil {
		in, out := &in.LogLevel, &out.LogLevel
		*out = new(string)
		**out = **in
	}
	if in.Confd != nil {
		in, out := &in.Confd, &out.Confd
		*out = new(ConfigDirSpec)
		**out = **in
	}
	if in.Checksd != nil {
		in, out := &in.Checksd, &out.Checksd
		*out = new(ConfigDirSpec)
		**out = **in
	}
	if in.PodLabelsAsTags != nil {
		in, out := &in.PodLabelsAsTags, &out.PodLabelsAsTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.PodAnnotationsAsTags != nil {
		in, out := &in.PodAnnotationsAsTags, &out.PodAnnotationsAsTags
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.CollectEvents != nil {
		in, out := &in.CollectEvents, &out.CollectEvents
		*out = new(bool)
		**out = **in
	}
	if in.LeaderElection != nil {
		in, out := &in.LeaderElection, &out.LeaderElection
		*out = new(bool)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.VolumeMounts != nil {
		in, out := &in.VolumeMounts, &out.VolumeMounts
		*out = make([]v1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Volumes != nil {
		in, out := &in.Volumes, &out.Volumes
		*out = make([]v1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.CriSocket != nil {
		in, out := &in.CriSocket, &out.CriSocket
		*out = new(CRISocketConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Dogstatsd != nil {
		in, out := &in.Dogstatsd, &out.Dogstatsd
		*out = new(DogstatsdConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeAgentConfig.
func (in *NodeAgentConfig) DeepCopy() *NodeAgentConfig {
	if in == nil {
		return nil
	}
	out := new(NodeAgentConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProcessSpec) DeepCopyInto(out *ProcessSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProcessSpec.
func (in *ProcessSpec) DeepCopy() *ProcessSpec {
	if in == nil {
		return nil
	}
	out := new(ProcessSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RbacConfig) DeepCopyInto(out *RbacConfig) {
	*out = *in
	if in.Create != nil {
		in, out := &in.Create, &out.Create
		*out = new(bool)
		**out = **in
	}
	if in.ServiceAccountName != nil {
		in, out := &in.ServiceAccountName, &out.ServiceAccountName
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RbacConfig.
func (in *RbacConfig) DeepCopy() *RbacConfig {
	if in == nil {
		return nil
	}
	out := new(RbacConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemProbeSpec) DeepCopyInto(out *SystemProbeSpec) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.ConntrackEnabled != nil {
		in, out := &in.ConntrackEnabled, &out.ConntrackEnabled
		*out = new(bool)
		**out = **in
	}
	if in.BPFDebugEnabled != nil {
		in, out := &in.BPFDebugEnabled, &out.BPFDebugEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.SecurityContext)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemProbeSpec.
func (in *SystemProbeSpec) DeepCopy() *SystemProbeSpec {
	if in == nil {
		return nil
	}
	out := new(SystemProbeSpec)
	in.DeepCopyInto(out)
	return out
}
