/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// OSPF summary address configuration to summarize external routes
type OspfSummaryAddressConfig struct {
	// Used to filter the advertisement of external routes into the OSPF domain. Setting this field to \"TRUE\" will enable the summarization of external routes that are covered by ip_prefix configuration. Setting this field to \"FALSE\" will filter the advertisement of external routes that are covered by ip_prefix configuration. 
	Advertise bool `json:"advertise,omitempty"`
	// OSPF Summary address in CIDR format
	Prefix string `json:"prefix"`
}
