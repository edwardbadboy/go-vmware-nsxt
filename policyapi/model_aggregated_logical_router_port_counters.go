/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type AggregatedLogicalRouterPortCounters struct {
	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	Rx *LogicalRouterPortCounters `json:"rx,omitempty"`
	Tx *LogicalRouterPortCounters `json:"tx,omitempty"`
}
