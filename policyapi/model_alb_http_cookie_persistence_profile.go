/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer HttpCookiePersistenceProfile object
type AlbHttpCookiePersistenceProfile struct {
	// If no persistence cookie was received from the client, always send it. Default value when not specified in API or module is interpreted by ALB Controller as false. 
	AlwaysSendCookie bool `json:"always_send_cookie,omitempty"`
	// HTTP cookie name for cookie persistence.
	CookieName string `json:"cookie_name,omitempty"`
	// Key name to use for cookie encryption.
	EncryptionKey string `json:"encryption_key,omitempty"`
	// Placeholder for description of property key of obj type HttpCookiePersistenceProfile field type str  type array. 
	Key []AlbHttpCookiePersistenceKey `json:"key,omitempty"`
	// The maximum lifetime of any session cookie. No value or 'zero' indicates no timeout. Allowed values are 1-14400. Special values are 0- 'No Timeout'. Unit is MIN. 
	Timeout int64 `json:"timeout,omitempty"`
}
