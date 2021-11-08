/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// DHCP server configuration. Please note, the realized-state of this entity returned by the \"GET /policy/api/v1/infra/realized-state/realized-entity\" with this entity policy-path is irrelevant with the application status of this entity. Please do not rely on this returned realized-state to determine how this dhcp-server-config was applied. The dhcp realization information was reflected in the realization states of the referencing Segment or T0/T1 gateway. 
type DhcpServerConfig struct {
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
	// Path of its parent
	ParentPath string `json:"parent_path,omitempty"`
	// Absolute path of this object
	Path string `json:"path,omitempty"`
	// This is a UUID generated by the system for realizing the entity object. In most cases this should be same as 'unique_id' of the entity. However, in some cases this can be different because of entities have migrated thier unique identifier to NSX Policy intent objects later in the timeline and did not use unique_id for realization. Realization id is helpful for users to debug data path to correlate the configuration with corresponding intent. 
	RealizationId string `json:"realization_id,omitempty"`
	// Path relative from its parent
	RelativePath string `json:"relative_path,omitempty"`
	// This is a UUID generated by the GM/LM to uniquely identify entites in a federated environment. For entities that are stretched across multiple sites, the same ID will be used on all the stretched sites. 
	UniqueId string `json:"unique_id,omitempty"`
	// subtree for this type within policy tree containing nested elements. 
	Children []ChildPolicyConfigResource `json:"children,omitempty"`
	// Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects. 
	MarkedForDelete bool `json:"marked_for_delete,omitempty"`
	// Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties. 
	Overridden bool `json:"overridden,omitempty"`
	// The reference to the edge cluster using the policy path of the edge cluster. Auto assigned if only one edge cluster is configured on enforcement-point. Modifying edge cluster will reallocate DHCP server to the new edge cluster. Please note that re-allocating edge-cluster will result in losing of all exisitng DHCP lease information. Change edge cluster only when losing DHCP leases is not a real problem, e.g. cross-site migration or failover and all client hosts will be reboot and get new IP addresses. 
	EdgeClusterPath string `json:"edge_cluster_path,omitempty"`
	// IP address lease time in seconds. 
	LeaseTime int64 `json:"lease_time,omitempty"`
	// Policy paths to edge nodes on which the DHCP servers run. The first edge node is assigned as active edge, and second one as stanby edge. If only one edge node is specified, the DHCP servers will run without HA support. When this property is not specified, edge nodes are auto-assigned during realization of the DHCP server. 
	PreferredEdgePaths []string `json:"preferred_edge_paths,omitempty"`
	// DHCP server address in CIDR format. Prefix length should be less than or equal to 30. DHCP server is deployed as DHCP relay service. This property is deprecated, use server_addresses instead. Both properties cannot be specified together with different new values. 
	ServerAddress string `json:"server_address,omitempty"`
	// DHCP server address in CIDR format. Both IPv4 and IPv6 address families are supported. Prefix length should be less than or equal to 30 for IPv4 address family and less than or equal to 126 for IPv6. When not specified, IPv4 value is auto-assigned to 100.96.0.1/30. Ignored when this object is configured at a Segment. 
	ServerAddresses []string `json:"server_addresses,omitempty"`
}
