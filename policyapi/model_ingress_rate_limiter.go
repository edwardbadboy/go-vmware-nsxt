/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// A shaper that specifies ingress rate properties in Mb/s
type IngressRateLimiter struct {
	Enabled bool `json:"enabled"`
	// Type rate limiter 
	ResourceType string `json:"resource_type"`
	// You can use the average bandwidth to reduce network congestion.
	AverageBandwidth int32 `json:"average_bandwidth,omitempty"`
	// The burst duration is set in the burst size setting.
	BurstSize int32 `json:"burst_size,omitempty"`
	// The peak bandwidth rate is used to support burst traffic.
	PeakBandwidth int32 `json:"peak_bandwidth,omitempty"`
}
