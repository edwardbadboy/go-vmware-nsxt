/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// The server will populate this field when returing the resource. Ignored on PUT and POST.
type SelfResourceLink struct {
	// Optional action
	Action string `json:"action,omitempty"`
	// Link to resource
	Href string `json:"href,omitempty"`
	// Custom relation type (follows RFC 5988 where appropriate definitions exist)
	Rel string `json:"rel,omitempty"`
}
