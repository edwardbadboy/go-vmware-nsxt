/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Collection of tags used in a policy group for a particular member type 
type GroupMemberTagsList struct {
	// Member type for which we will list the tags
	MemberType string `json:"member_type"`
	// List of tags for the member type
	Tags []string `json:"tags"`
}
