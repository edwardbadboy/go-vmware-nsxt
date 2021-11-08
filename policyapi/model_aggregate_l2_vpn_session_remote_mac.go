/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Aggregate of L2Vpn session remote mac across enforcement points. 
type AggregateL2VpnSessionRemoteMac struct {
	// Intent path of object, forward slashes must be escaped using %2F. 
	IntentPath string `json:"intent_path,omitempty"`
	// List of L2Vpn Session remote mac
	L2vpnSessionRemoteMacs []L2VpnSessionRemoteMacPerEp `json:"l2vpn_session_remote_macs,omitempty"`
}
