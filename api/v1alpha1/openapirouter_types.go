/*
Copyright 2025.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	gatewayv1 "sigs.k8s.io/gateway-api/apis/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// OpenApiRouterSpec defines the desired state of OpenApiRouter.
type OpenApiRouterSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	BackendRefs []gatewayv1.HTTPBackendRef `json:"backend,omitempty"`
}

// OpenApiRouterStatus defines the observed state of OpenApiRouter.
type OpenApiRouterStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// OpenApiRouter is the Schema for the openapirouters API.
type OpenApiRouter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpenApiRouterSpec   `json:"spec,omitempty"`
	Status OpenApiRouterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpenApiRouterList contains a list of OpenApiRouter.
type OpenApiRouterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpenApiRouter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OpenApiRouter{}, &OpenApiRouterList{})
}
