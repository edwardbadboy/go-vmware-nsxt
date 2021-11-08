/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer WafPSMMatchElement object
type AlbWafPsmMatchElement struct {
	// Mark this element excluded, like in '!ARGS password'. Default value when not specified in API or module is interpreted by ALB Controller as false. 
	Excluded bool `json:"excluded,omitempty"`
	// Match_element index.
	Index int64 `json:"index"`
	// The variable specification. For example ARGS or REQUEST_COOKIES. This can be a scalar like PATH_INFO. Enum options - WAF_VARIABLE_ARGS, WAF_VARIABLE_ARGS_GET, WAF_VARIABLE_ARGS_POST, WAF_VARIABLE_ARGS_NAMES, WAF_VARIABLE_REQUEST_COOKIES, WAF_VARIABLE_QUERY_STRING, WAF_VARIABLE_REQUEST_BASENAME, WAF_VARIABLE_REQUEST_URI, WAF_VARIABLE_PATH_INFO. 
	Name string `json:"name"`
	// The name of the request collection element. This can be empty, if we address the whole collection or a scalar element. 
	SubElement string `json:"sub_element,omitempty"`
}
