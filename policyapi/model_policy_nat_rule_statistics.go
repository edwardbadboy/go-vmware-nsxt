/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Gives the Statistics of a NAT rule. 
type PolicyNatRuleStatistics struct {
	// Gives the total number of active sessions. 
	ActiveSessions int64 `json:"active_sessions,omitempty"`
	// Gives the total number of bytes. 
	TotalBytes int64 `json:"total_bytes,omitempty"`
	// Gives the total number of packets. 
	TotalPackets int64 `json:"total_packets,omitempty"`
	// Timestamp when the data was last updated. 
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// The warning message about the NAT Rule Statistics. 
	WarningMessage string `json:"warning_message,omitempty"`
}
