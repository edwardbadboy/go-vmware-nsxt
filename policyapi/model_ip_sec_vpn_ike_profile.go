/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// IKE Profile is a reusable profile that captures IKE phase one negotiation parameters. Any changes affects all IPSec VPN sessions consuming this profile.
type IpSecVpnIkeProfile struct {
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
	// Diffie-Hellman group to be used if PFS is enabled. Default is GROUP14.
	DhGroups []string `json:"dh_groups,omitempty"`
	// Algorithm to be used for message digest during Internet Key Exchange(IKE) negotiation. A default value of SHA2_256 will be applied only when the supplied encryption algorithms contain either AES_128 or AES_256.
	DigestAlgorithms []string `json:"digest_algorithms,omitempty"`
	// Encryption algorithm is used during Internet Key Exchange(IKE) negotiation. Default is AES_128.
	EncryptionAlgorithms []string `json:"encryption_algorithms,omitempty"`
	// IKE protocol version to be used. IKE-Flex will initiate IKE-V2 and responds to both IKE-V1 and IKE-V2.
	IkeVersion string `json:"ike_version,omitempty"`
	// Life time for security association. Default is 86400 seconds (1 day).
	SaLifeTime int64 `json:"sa_life_time,omitempty"`
}
