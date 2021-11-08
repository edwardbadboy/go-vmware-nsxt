/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// IPSec VPN session status summary, gives total, failed, degraded and established IPSec VPN sessions.
type IPsecVpnikeSessionSummary struct {
	// Number of degraded sessions.
	DegradedSessions int64 `json:"degraded_sessions,omitempty"`
	// Number of established sessions.
	EstablishedSessions int64 `json:"established_sessions,omitempty"`
	// Number of failed sessions.
	FailedSessions int64 `json:"failed_sessions,omitempty"`
	// Total sessions configured.
	TotalSessions int64 `json:"total_sessions,omitempty"`
}
