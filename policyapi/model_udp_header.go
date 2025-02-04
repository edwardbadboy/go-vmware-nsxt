/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type UdpHeader struct {
	// Destination port of udp header
	DstPort int64 `json:"dst_port,omitempty"`
	// Source port of udp header
	SrcPort int64 `json:"src_port,omitempty"`
}
