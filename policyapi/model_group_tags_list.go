/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Collection of tags used in a policy group listed per member type 
type GroupTagsList struct {
	// Collection of tags used in a policy group listed per member type
	Results []GroupMemberTagsList `json:"results"`
}
