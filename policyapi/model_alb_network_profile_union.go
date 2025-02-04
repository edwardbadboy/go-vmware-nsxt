/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer NetworkProfileUnion object
type AlbNetworkProfileUnion struct {
	TcpFastPathProfile *AlbtcpFastPathProfile `json:"tcp_fast_path_profile,omitempty"`
	TcpProxyProfile *AlbtcpProxyProfile `json:"tcp_proxy_profile,omitempty"`
	// Configure one of either proxy or fast path profiles. Enum options - PROTOCOL_TYPE_TCP_PROXY, PROTOCOL_TYPE_TCP_FAST_PATH, PROTOCOL_TYPE_UDP_FAST_PATH, PROTOCOL_TYPE_UDP_PROXY. Allowed in Basic(Allowed values- PROTOCOL_TYPE_TCP_PROXY,PROTOCOL_TYPE_TCP_FAST_PATH,PROTOCOL_TYPE_UDP_FAST_PATH) edition, Essentials(Allowed values- PROTOCOL_TYPE_TCP_FAST_PATH,PROTOCOL_TYPE_UDP_FAST_PATH) edition, Enterprise edition. Default value when not specified in API or module is interpreted by ALB Controller as PROTOCOL_TYPE_TCP_PROXY. 
	Type_ string `json:"type"`
	UdpFastPathProfile *AlbudpFastPathProfile `json:"udp_fast_path_profile,omitempty"`
	UdpProxyProfile *AlbudpProxyProfile `json:"udp_proxy_profile,omitempty"`
}
