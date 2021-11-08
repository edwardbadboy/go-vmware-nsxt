/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Size of Directory Domain
type DirectoryDomainSize struct {
	// Number of groups
	GroupCount int32 `json:"group_count,omitempty"`
	// Number of group members
	GroupMemberCount int32 `json:"group_member_count,omitempty"`
	// Number of users
	UserCount int32 `json:"user_count,omitempty"`
}
