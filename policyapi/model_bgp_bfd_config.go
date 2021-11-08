/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// BFD configuration for BGP peers
type BgpBfdConfig struct {
	// Flag to enable BFD cofiguration.
	Enabled bool `json:"enabled,omitempty"`
	// Time interval between heartbeat packets in milliseconds. 
	Interval int32 `json:"interval,omitempty"`
	// Declare dead multiple. Number of times heartbeat packet is missed before BFD declares the neighbor is down. 
	Multiple int32 `json:"multiple,omitempty"`
}
