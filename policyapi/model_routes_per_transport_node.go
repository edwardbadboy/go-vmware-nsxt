/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// BGP routes per transport node.
type RoutesPerTransportNode struct {
	// Array of BGP neighbor route details for this transport node. 
	Routes []RouteDetails `json:"routes,omitempty"`
	// BGP neighbor source address.
	SourceAddress string `json:"source_address,omitempty"`
	// Transport node id
	TransportNodeId string `json:"transport_node_id,omitempty"`
}
