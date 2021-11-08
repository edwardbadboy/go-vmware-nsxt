/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Enables traffic limit for incoming/outgoing broadcast and multicast packets. Use 0 to disable rate limiting for a specific traffic type
type TrafficRateLimits struct {
	// Incoming broadcast traffic limit in packets per second
	RxBroadcast int32 `json:"rx_broadcast,omitempty"`
	// Incoming multicast traffic limit in packets per second
	RxMulticast int32 `json:"rx_multicast,omitempty"`
	// Outgoing broadcast traffic limit in packets per second
	TxBroadcast int32 `json:"tx_broadcast,omitempty"`
	// Outgoing multicast traffic limit in packets per second
	TxMulticast int32 `json:"tx_multicast,omitempty"`
}
