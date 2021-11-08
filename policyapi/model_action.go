/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Reaction Action is the action to take when the stipulated criteria specified in the event exist over the source. Some example actions include: - Notify Admin (or VMC's SRE) via email. - Populate a specific label with the IPSec VPN Session. - Remove the IPSec VPN Session from a specific label. 
type Action struct {
	// Reaction Action resource type. 
	ResourceType string `json:"resource_type"`
}
