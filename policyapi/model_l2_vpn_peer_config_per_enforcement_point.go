/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Peer config per Enforcement Point to configure the other side of the tunnel. 
type L2VpnPeerConfigPerEnforcementPoint struct {
	// Policy Path referencing the enforcement point to which the config belongs. 
	EnforcementPointPath string `json:"enforcement_point_path,omitempty"`
	ResourceType string `json:"resource_type"`
}
