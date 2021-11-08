/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Realized Firewall section
type RealizedFirewallSection struct {
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
	ResourceType string `json:"resource_type"`
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
	// Alarm info detail
	Alarms []PolicyAlarmResource `json:"alarms,omitempty"`
	// Desire state paths of this object
	IntentReference []string `json:"intent_reference,omitempty"`
	// Possible values could be UP, DOWN, UNKNOWN, FAILURE This list is not exhaustive. 
	OperationalStatus string `json:"operational_status,omitempty"`
	// It defines the root cause for operational status error. 
	OperationalStatusError string `json:"operational_status_error,omitempty"`
	// Possible values could be UP, DOWN, UNKNOWN, SUCCESS This list is not exhaustive. 
	PublishStatus string `json:"publish_status,omitempty"`
	// It defines the root cause for publish status error. 
	PublishStatusError string `json:"publish_status_error,omitempty"`
	// It defines error code for publish status error. 
	PublishStatusErrorCode int32 `json:"publish_status_error_code,omitempty"`
	// Error details for publish status. 
	PublishStatusErrorDetails []ConfigurationStateElement `json:"publish_status_error_details,omitempty"`
	// Realization API of this object on enforcement point
	RealizationApi string `json:"realization_api,omitempty"`
	// Realization id of this object
	RealizationSpecificIdentifier string `json:"realization_specific_identifier,omitempty"`
	// It define the root cause for runtime error. 
	RuntimeError string `json:"runtime_error,omitempty"`
	// Possible values could be UP, DOWN, UNKNOWN, DEGRADED This list is not exhaustive. 
	RuntimeStatus string `json:"runtime_status,omitempty"`
	// Realization state of this object
	State string `json:"state"`
	// Number of rules in this section.
	RuleCount int64 `json:"rule_count,omitempty"`
	// List of firewall rules in the section.
	Rules []RealizedFirewallRule `json:"rules,omitempty"`
	// Type of the rules which a section can contain.
	SectionType string `json:"section_type,omitempty"`
}
