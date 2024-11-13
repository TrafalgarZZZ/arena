// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AppStatusApplyConfiguration represents an declarative configuration of the AppStatus type for use
// with apply.
type AppStatusApplyConfiguration struct {
	HealthLastUpdateTime *v1.Time                                           `json:"healthLastUpdateTime,omitempty"`
	Deployments          map[string]ServeDeploymentStatusApplyConfiguration `json:"serveDeploymentStatuses,omitempty"`
	Status               *string                                            `json:"status,omitempty"`
	Message              *string                                            `json:"message,omitempty"`
}

// AppStatusApplyConfiguration constructs an declarative configuration of the AppStatus type for use with
// apply.
func AppStatus() *AppStatusApplyConfiguration {
	return &AppStatusApplyConfiguration{}
}

// WithHealthLastUpdateTime sets the HealthLastUpdateTime field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HealthLastUpdateTime field is set to the value of the last call.
func (b *AppStatusApplyConfiguration) WithHealthLastUpdateTime(value v1.Time) *AppStatusApplyConfiguration {
	b.HealthLastUpdateTime = &value
	return b
}

// WithDeployments puts the entries into the Deployments field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the Deployments field,
// overwriting an existing map entries in Deployments field with the same key.
func (b *AppStatusApplyConfiguration) WithDeployments(entries map[string]ServeDeploymentStatusApplyConfiguration) *AppStatusApplyConfiguration {
	if b.Deployments == nil && len(entries) > 0 {
		b.Deployments = make(map[string]ServeDeploymentStatusApplyConfiguration, len(entries))
	}
	for k, v := range entries {
		b.Deployments[k] = v
	}
	return b
}

// WithStatus sets the Status field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Status field is set to the value of the last call.
func (b *AppStatusApplyConfiguration) WithStatus(value string) *AppStatusApplyConfiguration {
	b.Status = &value
	return b
}

// WithMessage sets the Message field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Message field is set to the value of the last call.
func (b *AppStatusApplyConfiguration) WithMessage(value string) *AppStatusApplyConfiguration {
	b.Message = &value
	return b
}