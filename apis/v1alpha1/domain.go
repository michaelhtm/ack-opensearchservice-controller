// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	ackv1alpha1 "github.com/aws-controllers-k8s/runtime/apis/core/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DomainSpec defines the desired state of Domain.
type DomainSpec struct {

	// Options for all machine learning features for the specified domain.
	AIMLOptions *AIMLOptionsInput `json:"aimlOptions,omitempty"`
	// Identity and Access Management (IAM) policy document specifying the access
	// policies for the new domain.
	AccessPolicies *string `json:"accessPolicies,omitempty"`
	// Key-value pairs to specify advanced configuration options. The following
	// key-value pairs are supported:
	//
	//   - "rest.action.multi.allow_explicit_index": "true" | "false" - Note the
	//     use of a string rather than a boolean. Specifies whether explicit references
	//     to indexes are allowed inside the body of HTTP requests. If you want to
	//     configure access policies for domain sub-resources, such as specific indexes
	//     and domain APIs, you must disable this property. Default is true.
	//
	//   - "indices.fielddata.cache.size": "80" - Note the use of a string rather
	//     than a boolean. Specifies the percentage of heap space allocated to field
	//     data. Default is unbounded.
	//
	//   - "indices.query.bool.max_clause_count": "1024" - Note the use of a string
	//     rather than a boolean. Specifies the maximum number of clauses allowed
	//     in a Lucene boolean query. Default is 1,024. Queries with more than the
	//     permitted number of clauses result in a TooManyClauses error.
	//
	//   - "override_main_response_version": "true" | "false" - Note the use of
	//     a string rather than a boolean. Specifies whether the domain reports its
	//     version as 7.10 to allow Elasticsearch OSS clients and plugins to continue
	//     working with it. Default is false when creating a domain and true when
	//     upgrading a domain.
	//
	// For more information, see Advanced cluster parameters (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomain-configure-advanced-options).
	AdvancedOptions map[string]*string `json:"advancedOptions,omitempty"`
	// Options for fine-grained access control.
	AdvancedSecurityOptions *AdvancedSecurityOptionsInput `json:"advancedSecurityOptions,omitempty"`
	// Options for Auto-Tune.
	AutoTuneOptions *AutoTuneOptionsInput `json:"autoTuneOptions,omitempty"`
	// Container for the cluster configuration of a domain.
	ClusterConfig *ClusterConfig `json:"clusterConfig,omitempty"`
	// Key-value pairs to configure Amazon Cognito authentication. For more information,
	// see Configuring Amazon Cognito authentication for OpenSearch Dashboards (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/cognito-auth.html).
	CognitoOptions *CognitoOptions `json:"cognitoOptions,omitempty"`
	// Additional options for the domain endpoint, such as whether to require HTTPS
	// for all traffic.
	DomainEndpointOptions *DomainEndpointOptions `json:"domainEndpointOptions,omitempty"`
	// Container for the parameters required to enable EBS-based storage for an
	// OpenSearch Service domain.
	EBSOptions *EBSOptions `json:"ebsOptions,omitempty"`
	// Key-value pairs to enable encryption at rest.
	EncryptionAtRestOptions *EncryptionAtRestOptions `json:"encryptionAtRestOptions,omitempty"`
	// String of format Elasticsearch_X.Y or OpenSearch_X.Y to specify the engine
	// version for the OpenSearch Service domain. For example, OpenSearch_1.0 or
	// Elasticsearch_7.9. For more information, see Creating and managing Amazon
	// OpenSearch Service domains (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/createupdatedomains.html#createdomains).
	EngineVersion *string `json:"engineVersion,omitempty"`
	// Specify either dual stack or IPv4 as your IP address type. Dual stack allows
	// you to share domain resources across IPv4 and IPv6 address types, and is
	// the recommended option. If you set your IP address type to dual stack, you
	// can't change your address type later.
	IPAddressType *string `json:"ipAddressType,omitempty"`
	// Key-value pairs to configure log publishing.
	LogPublishingOptions map[string]*LogPublishingOption `json:"logPublishingOptions,omitempty"`
	// Name of the OpenSearch Service domain to create. Domain names are unique
	// across the domains owned by an account within an Amazon Web Services Region.
	// +kubebuilder:validation:Required
	Name *string `json:"name"`
	// Enables node-to-node encryption.
	NodeToNodeEncryptionOptions *NodeToNodeEncryptionOptions `json:"nodeToNodeEncryptionOptions,omitempty"`
	// Specifies a daily 10-hour time block during which OpenSearch Service can
	// perform configuration changes on the domain, including service software updates
	// and Auto-Tune enhancements that require a blue/green deployment. If no options
	// are specified, the default start time of 10:00 P.M. local time (for the Region
	// that the domain is created in) is used.
	OffPeakWindowOptions *OffPeakWindowOptions `json:"offPeakWindowOptions,omitempty"`
	// Software update options for the domain.
	SoftwareUpdateOptions *SoftwareUpdateOptions `json:"softwareUpdateOptions,omitempty"`
	// List of tags to add to the domain upon creation.
	Tags []*Tag `json:"tags,omitempty"`
	// Container for the values required to configure VPC access domains. If you
	// don't specify these values, OpenSearch Service creates the domain with a
	// public endpoint. For more information, see Launching your Amazon OpenSearch
	// Service domains using a VPC (https://docs.aws.amazon.com/opensearch-service/latest/developerguide/vpc.html).
	VPCOptions *VPCOptions `json:"vpcOptions,omitempty"`
}

// DomainStatus defines the observed state of Domain
type DomainStatus struct {
	// All CRs managed by ACK have a common `Status.ACKResourceMetadata` member
	// that is used to contain resource sync state, account ownership,
	// constructed ARN for the resource
	// +kubebuilder:validation:Optional
	ACKResourceMetadata *ackv1alpha1.ResourceMetadata `json:"ackResourceMetadata"`
	// All CRS managed by ACK have a common `Status.Conditions` member that
	// contains a collection of `ackv1alpha1.Condition` objects that describe
	// the various terminal states of the CR and its backend AWS service API
	// resource
	// +kubebuilder:validation:Optional
	Conditions []*ackv1alpha1.Condition `json:"conditions"`
	// Information about a configuration change happening on the domain.
	// +kubebuilder:validation:Optional
	ChangeProgressDetails *ChangeProgressDetails `json:"changeProgressDetails,omitempty"`
	// Creation status of an OpenSearch Service domain. True if domain creation
	// is complete. False if domain creation is still in progress.
	// +kubebuilder:validation:Optional
	Created *bool `json:"created,omitempty"`
	// Deletion status of an OpenSearch Service domain. True if domain deletion
	// is complete. False if domain deletion is still in progress. Once deletion
	// is complete, the status of the domain is no longer returned.
	// +kubebuilder:validation:Optional
	Deleted *bool `json:"deleted,omitempty"`
	// The dual stack hosted zone ID for the domain.
	// +kubebuilder:validation:Optional
	DomainEndpointV2HostedZoneID *string `json:"domainEndpointV2HostedZoneID,omitempty"`
	// Unique identifier for the domain.
	// +kubebuilder:validation:Optional
	DomainID *string `json:"domainID,omitempty"`
	// The status of any changes that are currently in progress for the domain.
	// +kubebuilder:validation:Optional
	DomainProcessingStatus *string `json:"domainProcessingStatus,omitempty"`
	// Domain-specific endpoint used to submit index, search, and data upload requests
	// to the domain.
	// +kubebuilder:validation:Optional
	Endpoint *string `json:"endpoint,omitempty"`
	// If IPAddressType to set to dualstack, a version 2 domain endpoint is provisioned.
	// This endpoint functions like a normal endpoint, except that it works with
	// both IPv4 and IPv6 IP addresses. Normal endpoints work only with IPv4 IP
	// addresses.
	// +kubebuilder:validation:Optional
	EndpointV2 *string `json:"endpointV2,omitempty"`
	// The key-value pair that exists if the OpenSearch Service domain uses VPC
	// endpoints. For example:
	//
	//    * IPv4 IP addresses - 'vpc','vpc-endpoint-h2dsd34efgyghrtguk5gt6j2foh4.us-east-1.es.amazonaws.com'
	//
	//    * Dual stack IP addresses - 'vpcv2':'vpc-endpoint-h2dsd34efgyghrtguk5gt6j2foh4.aos.us-east-1.on.aws'
	// +kubebuilder:validation:Optional
	Endpoints map[string]*string `json:"endpoints,omitempty"`
	// Information about the domain properties that are currently being modified.
	// +kubebuilder:validation:Optional
	ModifyingProperties []*ModifyingProperties `json:"modifyingProperties,omitempty"`
	// The status of the domain configuration. True if OpenSearch Service is processing
	// configuration changes. False if the configuration is active.
	// +kubebuilder:validation:Optional
	Processing *bool `json:"processing,omitempty"`
	// The current status of the domain's service software.
	// +kubebuilder:validation:Optional
	ServiceSoftwareOptions *ServiceSoftwareOptions `json:"serviceSoftwareOptions,omitempty"`
	// DEPRECATED. Container for parameters required to configure automated snapshots
	// of domain indexes.
	// +kubebuilder:validation:Optional
	SnapshotOptions *SnapshotOptions `json:"snapshotOptions,omitempty"`
	// The status of a domain version upgrade to a new version of OpenSearch or
	// Elasticsearch. True if OpenSearch Service is in the process of a version
	// upgrade. False if the configuration is active.
	// +kubebuilder:validation:Optional
	UpgradeProcessing *bool `json:"upgradeProcessing,omitempty"`
}

// Domain is the Schema for the Domains API
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
type Domain struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DomainSpec   `json:"spec,omitempty"`
	Status            DomainStatus `json:"status,omitempty"`
}

// DomainList contains a list of Domain
// +kubebuilder:object:root=true
type DomainList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Domain `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Domain{}, &DomainList{})
}
