/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer L4RuleMatchTarget object
type Albl4RuleMatchTarget struct {
	ClientIp *AlbIpAddrMatch `json:"client_ip,omitempty"`
	Port *Albl4RulePortMatch `json:"port,omitempty"`
	Protocol *Albl4RuleProtocolMatch `json:"protocol,omitempty"`
}
