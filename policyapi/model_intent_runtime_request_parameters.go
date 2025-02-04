/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Request parameters that represents a an intent path. 
type IntentRuntimeRequestParameters struct {
	// Policy Path referencing an intent object. 
	IntentPath string `json:"intent_path"`
	// Policy Path referencing a site. This is applicable only on a GlobalManager. If no site_path is specified, then based on the span of the intent the response will be fetched from the respective sites 
	SitePath string `json:"site_path,omitempty"`
}
