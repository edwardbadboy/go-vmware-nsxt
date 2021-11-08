/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Global configuration
type FipsGlobalConfig struct {
	// When this flag is set to true FIPS mode will be set on ssl encryptions of load balancer feature.
	LbFipsEnabled bool `json:"lb_fips_enabled,omitempty"`
	// When this flag is set to true FIPS mode will be set on ssl encryptions of TLS inspection feature.
	TlsFipsEnabled bool `json:"tls_fips_enabled,omitempty"`
}
