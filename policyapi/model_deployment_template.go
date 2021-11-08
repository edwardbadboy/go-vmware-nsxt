/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Deployment Template holds the attributes specific to partner for which the service is created. These attributes are opaque to NSX.
type DeploymentTemplate struct {
	// List of attributes specific to a partner for which the service is created. There attributes are passed on to the partner appliance and is opaque to the NSX Manager.
	Attributes []Attribute `json:"attributes,omitempty"`
	// Deployment Template name.
	Name string `json:"name,omitempty"`
}
