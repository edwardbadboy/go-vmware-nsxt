/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Parameters to let the admin specify a relative position of a security policy or rule w.r.t to another one. 
type PolicyInsertParameters struct {
	// The security policy/rule path if operation is 'insert_after' or 'insert_before' 
	AnchorPath string `json:"anchor_path,omitempty"`
	// Operation
	Operation string `json:"operation,omitempty"`
}
