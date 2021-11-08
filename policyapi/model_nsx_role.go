/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Role
type NsxRole struct {
	// Please use the /user-info/permissions api to get the permission that the user has on each feature.
	Permissions []string `json:"permissions,omitempty"`
	// This field represents the identifier of the role. With the introduction of custom roles, this field is no longer an enum.
	Role string `json:"role"`
}
