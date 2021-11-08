/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type Ipv4Header struct {
	// The destination ip address.
	DstIp string `json:"dst_ip,omitempty"`
	// IP flags
	Flags int64 `json:"flags,omitempty"`
	// IP protocol - defaults to ICMP
	Protocol int64 `json:"protocol,omitempty"`
	// The source ip address.
	SrcIp string `json:"src_ip,omitempty"`
	// This is used together with src_ip to calculate dst_ip for broadcast when dst_ip is not given; not used in all other cases.
	SrcSubnetPrefixLen int64 `json:"src_subnet_prefix_len,omitempty"`
	// Time to live (ttl)
	Ttl int64 `json:"ttl,omitempty"`
}
