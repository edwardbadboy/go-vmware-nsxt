/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// PIM (Protocol Independent Multicast) RP (Randezvous Point) mapping details. 
type PimRpMapping struct {
	// Multicast group address.
	Group string `json:"group,omitempty"`
	// Value of this field will be true if this edge transport node acts as rendezvous point, otherwise false. 
	IsRp bool `json:"is_rp,omitempty"`
	// Outgoing/Egress interface for multicast traffic.
	OutgoingInterface string `json:"outgoing_interface,omitempty"`
	// RP (Randezvous Point) address.
	RpAddress string `json:"rp_address,omitempty"`
	// Source of learning RP information. Either Static RP configured or RP learned via BSR (Bootstrap Router). 
	Source string `json:"source,omitempty"`
}
