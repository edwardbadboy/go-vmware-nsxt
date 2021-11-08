/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Cloud Native Common Agent configuration that can be applied to host nodes. 
type CommonAgentHostConfigurationInfo struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected.
	Revision int32 `json:"_revision,omitempty"`
	// Timestamp of resource creation
	CreateTime int64 `json:"_create_time,omitempty"`
	// ID of the user who created this resource
	CreateUser string `json:"_create_user,omitempty"`
	// Timestamp of last modification
	LastModifiedTime int64 `json:"_last_modified_time,omitempty"`
	// ID of the user who last modified this resource
	LastModifiedUser string `json:"_last_modified_user,omitempty"`
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this           entity. 
	Protection string `json:"_protection,omitempty"`
	// Indicates system owned resource
	SystemOwned bool `json:"_system_owned,omitempty"`
	// Description of this resource
	Description string `json:"description,omitempty"`
	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`
	// Unique identifier of this resource
	Id string `json:"id,omitempty"`
	// The type of this resource.
	ResourceType string `json:"resource_type,omitempty"`
	// Opaque identifiers meaningful to the API user
	Tags []Tag `json:"tags,omitempty"`
	// A ingress certificate to verify the identity of brokers. 
	IngressCertificate string `json:"ingress_certificate,omitempty"`
	// List of Cloud Native Platform ingress endpoints that host nodes contact initially. 
	IngressEndpoint []CommonAgentEndpointInfo `json:"ingress_endpoint,omitempty"`
	// A Kafka broker certificate to verify the identity of brokers. 
	KafkaCertificate string `json:"kafka_certificate,omitempty"`
	// List of Cloud Native Platform broker endpoints that host nodes contact initially. 
	KafkaEndpoint []CommonAgentEndpointInfo `json:"kafka_endpoint,omitempty"`
	// Cluster id of the NSX Manager cluster. 
	NsxClusterId string `json:"nsx_cluster_id,omitempty"`
	// List of private IP prefix that Cloud Native Common Agent network flow is collected from. 
	PrivateIpRange []CommonAgentPrivateIpRangeInfo `json:"private_ip_range,omitempty"`
	// A truststore to establish the trust between NSX and Cloud Native Platform. 
	Truststore string `json:"truststore,omitempty"`
}
