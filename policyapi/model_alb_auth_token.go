/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// ALB Auth Token
type AlbAuthToken struct {
	// Expiry time of the token will be set by LCM at the time of Enforcement Point Creation. 
	ExpiresAt string `json:"expires_at,omitempty"`
	// Hours to validate the token 
	Hours string `json:"hours"`
	// Token for Avi Controller.
	Token string `json:"token,omitempty"`
	// controller username. 
	Username string `json:"username"`
}
