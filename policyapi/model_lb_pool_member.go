/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Pool member.
type LbPoolMember struct {
	// Member admin state.
	AdminState string `json:"admin_state,omitempty"`
	// Backup servers are typically configured with a sorry page indicating to the user that the application is currently unavailable. While the pool is active (a specified minimum number of pool members are active) BACKUP members are skipped during server selection. When the pool is inactive, incoming connections are sent to only the BACKUP member(s). 
	BackupMember bool `json:"backup_member,omitempty"`
	// Pool member name.
	DisplayName string `json:"display_name,omitempty"`
	// Pool member IP address.
	IpAddress string `json:"ip_address"`
	// To ensure members are not overloaded, connections to a member can be capped by the load balancer. When a member reaches this limit, it is skipped during server selection. If it is not specified, it means that connections are unlimited. 
	MaxConcurrentConnections int64 `json:"max_concurrent_connections,omitempty"`
	// If port is specified, all connections will be sent to this port. Only single port is supported. If unset, the same port the client connected to will be used, it could be overrode by default_pool_member_port setting in virtual server. The port should not specified for port range case. 
	Port string `json:"port,omitempty"`
	// Pool member weight is used for WEIGHTED_ROUND_ROBIN balancing algorithm. The weight value would be ignored in other algorithms. 
	Weight int64 `json:"weight,omitempty"`
}
