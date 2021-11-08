/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// The setting is used to add, update or remove pool members from pool. For static pool members, admin_state, display_name and weight can be updated. For dynamic pool members, only admin_state can be updated. 
type PoolMemberSetting struct {
	// Member admin state
	AdminState string `json:"admin_state,omitempty"`
	// Only applicable to static pool members. If supplied for a pool defined by a grouping object, update API would fail. 
	DisplayName string `json:"display_name,omitempty"`
	// Pool member IP address
	IpAddress string `json:"ip_address"`
	// Pool member port number
	Port string `json:"port,omitempty"`
	// Only applicable to static pool members. If supplied for a pool defined by a grouping object, update API would fail. 
	Weight int64 `json:"weight,omitempty"`
}
