/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer CustomParams object
type AlbCustomParams struct {
	// Placeholder for description of property is_dynamic of obj type CustomParams field type str  type boolean. Default value when not specified in API or module is interpreted by ALB Controller as false. 
	IsDynamic bool `json:"is_dynamic,omitempty"`
	// Placeholder for description of property is_sensitive of obj type CustomParams field type str  type boolean. Default value when not specified in API or module is interpreted by ALB Controller as false. 
	IsSensitive bool `json:"is_sensitive,omitempty"`
	// Name of the object.
	Name string `json:"name"`
	// value of CustomParams.
	Value string `json:"value,omitempty"`
}
