// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TokenInitParameters struct {

	// (String) The API token creation date.
	// The API token creation date.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// (String) The environment the token will have access to. Use "*" for all environments. By default, it will have access to all environments.
	// The environment the token will have access to. Use `"*"` for all environments. By default, it will have access to all environments.
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// (String) The API token expiration date.
	// The API token expiration date.
	ExpiresAt *string `json:"expiresAt,omitempty" tf:"expires_at,omitempty"`

	// (Set of String) The project(s) the token will have access to. Use ["*"] for all projects. By default, it will have access to all projects.
	// The project(s) the token will have access to. Use `["*"]` for all projects. By default, it will have access to all projects.
	// +listType=set
	Projects []*string `json:"projects,omitempty" tf:"projects,omitempty"`

	// (String, Sensitive) The API token secret.
	// The API token secret.
	SecretSecretRef *v1.SecretKeySelector `json:"secretSecretRef,omitempty" tf:"-"`

	// (String) The type of the API token. Can be client, admin or frontend
	// The type of the API token. Can be `client`, `admin` or `frontend`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String)
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type TokenObservation struct {

	// (String) The API token creation date.
	// The API token creation date.
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// (String) The environment the token will have access to. Use "*" for all environments. By default, it will have access to all environments.
	// The environment the token will have access to. Use `"*"` for all environments. By default, it will have access to all environments.
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// (String) The API token expiration date.
	// The API token expiration date.
	ExpiresAt *string `json:"expiresAt,omitempty" tf:"expires_at,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Set of String) The project(s) the token will have access to. Use ["*"] for all projects. By default, it will have access to all projects.
	// The project(s) the token will have access to. Use `["*"]` for all projects. By default, it will have access to all projects.
	// +listType=set
	Projects []*string `json:"projects,omitempty" tf:"projects,omitempty"`

	// (String) The type of the API token. Can be client, admin or frontend
	// The type of the API token. Can be `client`, `admin` or `frontend`
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String)
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type TokenParameters struct {

	// (String) The API token creation date.
	// The API token creation date.
	// +kubebuilder:validation:Optional
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// (String) The environment the token will have access to. Use "*" for all environments. By default, it will have access to all environments.
	// The environment the token will have access to. Use `"*"` for all environments. By default, it will have access to all environments.
	// +kubebuilder:validation:Optional
	Environment *string `json:"environment,omitempty" tf:"environment,omitempty"`

	// (String) The API token expiration date.
	// The API token expiration date.
	// +kubebuilder:validation:Optional
	ExpiresAt *string `json:"expiresAt,omitempty" tf:"expires_at,omitempty"`

	// (Set of String) The project(s) the token will have access to. Use ["*"] for all projects. By default, it will have access to all projects.
	// The project(s) the token will have access to. Use `["*"]` for all projects. By default, it will have access to all projects.
	// +kubebuilder:validation:Optional
	// +listType=set
	Projects []*string `json:"projects,omitempty" tf:"projects,omitempty"`

	// (String, Sensitive) The API token secret.
	// The API token secret.
	// +kubebuilder:validation:Optional
	SecretSecretRef *v1.SecretKeySelector `json:"secretSecretRef,omitempty" tf:"-"`

	// (String) The type of the API token. Can be client, admin or frontend
	// The type of the API token. Can be `client`, `admin` or `frontend`
	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// (String)
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// TokenSpec defines the desired state of Token
type TokenSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TokenParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider TokenInitParameters `json:"initProvider,omitempty"`
}

// TokenStatus defines the observed state of Token.
type TokenStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TokenObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Token is the Schema for the Tokens API. Provides a resource for managing unleash api tokens.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,unleash}
type Token struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.type) || (has(self.initProvider) && has(self.initProvider.type))",message="spec.forProvider.type is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.username) || (has(self.initProvider) && has(self.initProvider.username))",message="spec.forProvider.username is a required parameter"
	Spec   TokenSpec   `json:"spec"`
	Status TokenStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TokenList contains a list of Tokens
type TokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Token `json:"items"`
}

// Repository type metadata.
var (
	Token_Kind             = "Token"
	Token_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Token_Kind}.String()
	Token_KindAPIVersion   = Token_Kind + "." + CRDGroupVersion.String()
	Token_GroupVersionKind = CRDGroupVersion.WithKind(Token_Kind)
)

func init() {
	SchemeBuilder.Register(&Token{}, &TokenList{})
}
