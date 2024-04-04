// Code generated by applyconfiguration-gen. DO NOT EDIT.

package applyconfig

import (
	v1 "github.com/applygen-cycle-repro/applyconfig/resource.io/v1"
	resourceiov1 "github.io/applygen-cycle-repro/applyconfig/resource.io/v1"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
)

// ForKind returns an apply configuration type for the given GroupVersionKind, or nil if no
// apply configuration type exists for the given GroupVersionKind.
func ForKind(kind schema.GroupVersionKind) interface{} {
	switch kind {
	// Group=resource.io, Version=v1
	case v1.SchemeGroupVersion.WithKind("Priority"):
		return &resourceiov1.PriorityApplyConfiguration{}
	case v1.SchemeGroupVersion.WithKind("SubPriority"):
		return &resourceiov1.SubPriorityApplyConfiguration{}

	}
	return nil
}