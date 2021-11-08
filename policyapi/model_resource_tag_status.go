/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// It represents tag operation status for a resource and details of the failure if any. 
type ResourceTagStatus struct {
	// Details about the error if any
	Details string `json:"details,omitempty"`
	// Resource display name
	ResourceDisplayName string `json:"resource_display_name,omitempty"`
	// Resource id
	ResourceId string `json:"resource_id"`
	// Status of tag apply or remove operation
	TagStatus string `json:"tag_status"`
}
