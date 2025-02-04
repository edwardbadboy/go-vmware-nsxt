/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type IgmpMembershipCsvRecord struct {
	// Multicast group address.
	Group string `json:"group,omitempty"`
	// IGMP version.
	IgmpVersion int64 `json:"igmp_version,omitempty"`
	// Interface on which multicast group membership is learned. 
	Interface_ string `json:"interface,omitempty"`
	// Number of multicast sources.
	NoOfSources string `json:"no_of_sources,omitempty"`
	// IP address of multicast source.
	SourceAddress string `json:"source_address,omitempty"`
	// Transport node uuid or policy path.
	TransportNode string `json:"transport_node,omitempty"`
	// Multicast group membership active time.
	Uptime string `json:"uptime,omitempty"`
}
