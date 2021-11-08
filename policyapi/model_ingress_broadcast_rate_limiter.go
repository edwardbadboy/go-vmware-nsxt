/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// A shaper that specifies ingress rate properties in kb/s
type IngressBroadcastRateLimiter struct {
	Enabled bool `json:"enabled"`
	// Type rate limiter 
	ResourceType string `json:"resource_type"`
	// Average bandwidth in kb/s
	AverageBandwidth int32 `json:"average_bandwidth,omitempty"`
	// Burst size in bytes
	BurstSize int32 `json:"burst_size,omitempty"`
	// Peak bandwidth in kb/s
	PeakBandwidth int32 `json:"peak_bandwidth,omitempty"`
}
