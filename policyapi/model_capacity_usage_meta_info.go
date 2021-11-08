/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type CapacityUsageMetaInfo struct {
	// Timestamp at which capacity usage was last calculated
	LastUpdatedTimestamp int64 `json:"last_updated_timestamp"`
	// Indicates the maximum global threshold percentage 
	MaxGlobalThresholdPercentage float32 `json:"max_global_threshold_percentage"`
	// Indicates the minimum global threshold percentage 
	MinGlobalThresholdPercentage float32 `json:"min_global_threshold_percentage"`
}
