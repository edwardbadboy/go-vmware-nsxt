/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Distributed virtual portgroup on a VC
type DistributedVirtualPortgroup struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	Self *SelfResourceLink `json:"_self,omitempty"`
	// Timestamp of last modification
	LastSyncTime int64 `json:"_last_sync_time,omitempty"`
	// Description of this resource
	Description string `json:"description,omitempty"`
	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`
	// The type of this resource.
	ResourceType string `json:"resource_type"`
	// Specifies list of scope of discovered resource. e.g. if VHC path is associated with principal identity, who owns the discovered resource, then scope id will be VHC path and scope type will be VHC. 
	Scope []DiscoveredResourceScope `json:"scope,omitempty"`
	// Opaque identifiers meaningful to the API user
	Tags []Tag `json:"tags,omitempty"`
	// Id of the portgroup, eg. a mo-ref from VC.
	CmLocalId string `json:"cm_local_id,omitempty"`
	// External id of the virtual portgroup
	ExternalId string `json:"external_id,omitempty"`
	// Portgroup type like DistributedVirtualPortgroup
	OriginType string `json:"origin_type,omitempty"`
	// For distributed virtual portgroup, backing type is standard. For logical switch portgroup, the backing type is set to nsx. 
	BackingType string `json:"backing_type,omitempty"`
	// Generated UUID of the portgroup
	Key string `json:"key,omitempty"`
	// This parameters reflects the managed entity status of the portgroup as reported by VC. 
	OverallStatus string `json:"overall_status,omitempty"`
}
