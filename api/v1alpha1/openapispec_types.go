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
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// OpenApiSpecRef defines a reference to an OpenApiSpec.
type OpenApiSpecRef struct {
	// Name is the name of the OpenApiSpec.
	Name string `json:"name,omitempty"`

	// Namespace is the namespace of the OpenApiSpec.
	// +optional
	Namespace string `json:"namespace,omitempty"`
}

// OpenApiSpecSpec defines the desired state of OpenApiSpec.
type OpenApiSpecSpec struct {
	Document string `json:"document,omitempty"`
	// TODO implement other ways of specifying the document, such as remotely or from a ConfigMap
}

// OpenApiSpecStatus defines the observed state of OpenApiSpec.
type OpenApiSpecStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// OpenApiSpec is the Schema for the openapispecs API.
type OpenApiSpec struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OpenApiSpecSpec   `json:"spec,omitempty"`
	Status OpenApiSpecStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OpenApiSpecList contains a list of OpenApiSpec.
type OpenApiSpecList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OpenApiSpec `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OpenApiSpec{}, &OpenApiSpecList{})
}

// TODO implement validating webhook on OpenApiSpec to ensure that the spec.document is valid
