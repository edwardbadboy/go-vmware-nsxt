/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Status of the Identity Firewall Compute Collection's transport node. 
type IdfwTransportNodeCondition struct {
	// Transport node status for IDFW compute collection.
	Status string `json:"status"`
	// IDFW Compute collection's transport node condition.
	StatusDetail string `json:"status_detail,omitempty"`
}
