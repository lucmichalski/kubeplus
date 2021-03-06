package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ResourceComposition is specification for a ResourceComposition resource
// +k8s:openapi-gen=true
type ResourceComposition struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourceCompositionSpec   `json:"spec"`
	Status ResourceCompositionStatus `json:"status"`
}

// ResourceCompositionSpec is the spec for a ResourceComposition resource
// +k8s:openapi-gen=true
type ResourceCompositionSpec struct {
	// LabelSelector that selects resources of that label
	//LabelSelector map[string]string `json:"labelSelector,omitempty"` 
	// List of stack elements that forms this Platform Workflow
	//StackElements []StackElements `json:"stackElements"`
	// Name of CRD to register
	NewResource NewResource `json:"newResource"`
}

type StackElements struct {
	// Name of Kubernetes Kind - could be Built-in Kind or Custom Kind
	Kind string `json:"kind"`
	// Name of the Resource
	Name string `json:"name"`
	// Namespace of the Resource
	// If not specified then 'default' Namespace is assumed
	// +optional
	Namespace string `json:"namespace"`
	// +optional
	DependsOn []DependsOn `json:"dependsOn,omitempty"`
}

type DependsOn struct {
	// Name of the Resource that the Resource in which dependsOn is included actually dependsOn
	Name string `json:"name"`
}

type NewResource struct {
	Resource Res `json:"resource"`
	// Helm chart URL
	ChartURL string `json:"chartURL"`
	// Chart name
	ChartName string `json:"chartName"`
	// Values
	//Values []Values `json:"values,omitempty"`
}

type Res struct {
	// Kind of the Custom API
	Kind string `json:"kind"`
	// Version of the API Custom API
	Version string `json:"version"`
	// Group of the Custom API
	Group string `json:"group"`
	// Plural name for the Custom API
	Plural string `json:"plural"`
}

type Values struct {
	// Name - as used in values.yaml
	Name string `json:"name"`
	// Value - as specified in values.yaml
	Value string `json:"value"`
}

// ResourceCompositionStatus is the status for a ResourceComposition resource.
// +k8s:openapi-gen=true
type ResourceCompositionStatus struct {
	Status             string   `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ResourceCompositionList is a list of ResourceComposition resources
type ResourceCompositionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ResourceComposition `json:"items"`
}


// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// ResourcePolicy is specification for a ResourcePolicy resource
type ResourcePolicy struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ResourcePolicySpec   `json:"spec"`
	Status ResourcePolicyStatus `json:"status"`
}

type ResourcePolicyStatus struct {
	Status             string   `json:"status"`
}

type ResourcePolicySpec struct {
	Resource Res `json:"resource"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// ResourcePolicyList is a list of ResourcePolicy resources
type ResourcePolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ResourcePolicy `json:"items"`
}

