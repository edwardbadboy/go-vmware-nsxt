/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type BinaryPacketData struct {
	// If the requested frame_size is too small (given the payload and traceflow metadata requirement of 16 bytes), the traceflow request will fail with an appropriate message.  The frame will be zero padded to the requested size.
	FrameSize int64 `json:"frame_size,omitempty"`
	// Packet configuration
	ResourceType string `json:"resource_type"`
	// A flag, when set true, indicates that the traceflow packet is of L3 routing.
	Routed bool `json:"routed,omitempty"`
	// This type takes effect only for IP packet.
	TransportType string `json:"transport_type,omitempty"`
	// Up to 1000 bytes of payload may be supplied (with a base64-encoded length of 1336 bytes.) Additional bytes of traceflow metadata will be appended to the payload. The payload must contain all headers (Ethernet, IP, etc). Note that VLAN is not supported in the logical space. Hence, payload must not contain 802.1Q headers.
	Payload string `json:"payload,omitempty"`
}
