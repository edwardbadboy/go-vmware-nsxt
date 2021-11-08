/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer HTTPRedirectAction object
type AlbhttpRedirectAction struct {
	Host *AlburiParam `json:"host,omitempty"`
	// Keep or drop the query of the incoming request URI in the redirected URI. Default value when not specified in API or module is interpreted by ALB Controller as true. 
	KeepQuery bool `json:"keep_query,omitempty"`
	Path *AlburiParam `json:"path,omitempty"`
	// Port to which redirect the request. Allowed values are 1-65535. 
	Port int64 `json:"port,omitempty"`
	// Protocol type. Enum options - HTTP, HTTPS. 
	Protocol string `json:"protocol"`
	// HTTP redirect status code. Enum options - HTTP_REDIRECT_STATUS_CODE_301, HTTP_REDIRECT_STATUS_CODE_302, HTTP_REDIRECT_STATUS_CODE_307. Default value when not specified in API or module is interpreted by ALB Controller as HTTP_REDIRECT_STATUS_CODE_302. 
	StatusCode string `json:"status_code,omitempty"`
}
