/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// New Role
type NewRole struct {
	// New role description
	NewRoleDescription string `json:"new_role_description,omitempty"`
	// New role id
	NewRoleId string `json:"new_role_id"`
	// New role name
	NewRoleName string `json:"new_role_name"`
}
