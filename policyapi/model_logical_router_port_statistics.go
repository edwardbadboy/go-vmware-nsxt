/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type LogicalRouterPortStatistics struct {
	// The ID of the logical router port
	LogicalRouterPortId string `json:"logical_router_port_id"`
	// Per Node Statistics
	PerNodeStatistics []LogicalRouterPortStatisticsPerNode `json:"per_node_statistics,omitempty"`
}
