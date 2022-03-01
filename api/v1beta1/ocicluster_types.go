/*
Copyright (c) 2021, 2022 Oracle and/or its affiliates.

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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	clusterv1 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
// Important: Run "make" to regenerate code after modifying this file

const (
	// ClusterFinalizer allows OCIClusterReconciler to clean up OCI resources associated with OCICluster before
	// removing it from the apiserver.
	ClusterFinalizer = "ocicluster.infrastructure.cluster.x-k8s.io"
)

// OCIClusterSpec defines the desired state of OciCluster
type OCIClusterSpec struct {
	// NetworkSpec encapsulates all things related to OCI network.
	// +optional
	NetworkSpec NetworkSpec `json:"networkSpec,omitempty"`
	// Free-form tags for this resource
	// +optional
	FreeformTags map[string]string `json:"freeformTags,omitempty"`

	// Defined tags for this resource. Each key is predefined and scoped to a
	// namespace. For more information, see Resource Tags (https://docs.cloud.oracle.com/iaas/Content/General/Concepts/resourcetags.htm).
	// Example: `{"Operations": {"CostCenter": "42"}}`
	// +optional
	DefinedTags map[string]map[string]string `json:"definedTags,omitempty"`

	// Compartment to create the cluster network.
	CompartmentId string `mandatory:"true" json:"compartmentId"`

	// ControlPlaneEndpoint represents the endpoint used to communicate with the control plane.
	// +optional
	ControlPlaneEndpoint clusterv1.APIEndpoint `json:"controlPlaneEndpoint"`
}

// OCIClusterStatus defines the observed state of OCICluster
type OCIClusterStatus struct {
	// +optional
	FailureDomains clusterv1.FailureDomains `json:"failureDomains,omitempty"`

	// +optional
	Ready bool `json:"ready"`
	// NetworkSpec encapsulates all things related to OCI network.
	// +optional
	Conditions clusterv1.Conditions `json:"conditions,omitempty"`
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// OCICluster is the Schema for the ociclusters API
type OCICluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OCIClusterSpec   `json:"spec,omitempty"`
	Status OCIClusterStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// OCIClusterList contains a list of OciCluster
type OCIClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OCICluster `json:"items"`
}

// GetConditions returns the list of conditions for an OCICluster API object.
func (c *OCICluster) GetConditions() clusterv1.Conditions {
	return c.Status.Conditions
}

// SetConditions will set the given conditions on an OCICluster object.
func (c *OCICluster) SetConditions(conditions clusterv1.Conditions) {
	c.Status.Conditions = conditions
}

func init() {
	SchemeBuilder.Register(&OCICluster{}, &OCIClusterList{})
}