/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer HealthMonitorAuthInfo object
type AlbHealthMonitorAuthInfo struct {
	// Password for server authentication.
	Password string `json:"password,omitempty"`
	// Username for server authentication.
	Username string `json:"username,omitempty"`
}
