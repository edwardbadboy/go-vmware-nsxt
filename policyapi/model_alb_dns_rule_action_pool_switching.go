/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer DnsRuleActionPoolSwitching object
type AlbDnsRuleActionPoolSwitching struct {
	// Reference of the pool group to serve the passthrough DNS query which cannot be served locally. It is a reference to an object of type PoolGroup. 
	PoolGroupPath string `json:"pool_group_path,omitempty"`
	// Reference of the pool to serve the passthrough DNS query which cannot be served locally. It is a reference to an object of type Pool. 
	PoolPath string `json:"pool_path,omitempty"`
}
