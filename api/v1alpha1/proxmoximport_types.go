package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ProxmoxImportSpec defines the desired state of ProxmoxImport
type ProxmoxImportSpec struct {
	Url string `json:"url,omitempty"`
}

// ProxmoxImportStatus defines the observed state of ProxmoxImport
type ProxmoxImportStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// ProxmoxImport is the Schema for the proxmoximports API
type ProxmoxImport struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ProxmoxImportSpec   `json:"spec,omitempty"`
	Status ProxmoxImportStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// ProxmoxImportList contains a list of ProxmoxImport
type ProxmoxImportList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ProxmoxImport `json:"items"`
}

func init() {
	objectTypes = append(objectTypes, &ProxmoxImport{}, &ProxmoxImportList{})
}
