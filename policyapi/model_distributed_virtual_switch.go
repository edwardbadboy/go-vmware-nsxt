/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// DistributedVirtualSwitch on a VC
type DistributedVirtualSwitch struct {
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
	// ID of the virtual switch in compute manager
	CmLocalId string `json:"cm_local_id,omitempty"`
	// External id of the virtual switch
	ExternalId string `json:"external_id,omitempty"`
	// ID of the compute manager where this virtual switch is discovered. 
	OriginId string `json:"origin_id,omitempty"`
	// Switch type like VmwareDistributedVirtualSwitch
	OriginType string `json:"origin_type,omitempty"`
	// Array of discovered nodes connected to this switch.
	DiscoveredNodes []DiscoveredNode `json:"discovered_nodes,omitempty"`
	// It contains information about VMware specific multiple dynamic LACP groups. 
	LacpGroupConfigs []LacpGroupConfigInfo `json:"lacp_group_configs,omitempty"`
	// Key-Value map of additional properties of switch
	OriginProperties []KeyValuePair `json:"origin_properties,omitempty"`
	// The uniform name of uplink ports on each host.
	UplinkPortNames []string `json:"uplink_port_names,omitempty"`
	UplinkPortgroup *DistributedVirtualPortgroup `json:"uplink_portgroup,omitempty"`
	// UUID of the switch
	Uuid string `json:"uuid,omitempty"`
}
