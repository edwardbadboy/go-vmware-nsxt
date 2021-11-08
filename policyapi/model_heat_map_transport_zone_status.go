/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type HeatMapTransportZoneStatus struct {
	// Number of transport nodes that are degraded
	DegradedCount int32 `json:"degraded_count,omitempty"`
	// Number of transport nodes that are down
	DownCount int32 `json:"down_count,omitempty"`
	// Number of transport nodes with unknown status
	UnknownCount int32 `json:"unknown_count,omitempty"`
	// Number of transport nodes that are up
	UpCount int32 `json:"up_count,omitempty"`
}
