/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Contains IPv6 related discovery options.
type IPv6DiscoveryOptions struct {
	// Enable this method will snoop the DHCPv6 message transaction which a VM makes with a DHCPv6 server. From the transaction, we learn the IPv6 addresses assigned by the DHCPv6 server to this VM along with its lease time. 
	DhcpSnoopingV6Enabled bool `json:"dhcp_snooping_v6_enabled,omitempty"`
	NdSnoopingConfig *NdSnoopingConfig `json:"nd_snooping_config,omitempty"`
	// Enable this method will learn the IPv6 addresses which are configured on interfaces of a VM with the help of the VMTools software. 
	VmtoolsV6Enabled bool `json:"vmtools_v6_enabled,omitempty"`
}
