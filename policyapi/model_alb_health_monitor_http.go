/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer HealthMonitorHttp object
type AlbHealthMonitorHttp struct {
	// Type of the authentication method. Enum options - AUTH_BASIC, AUTH_NTLM. Allowed in Basic edition, Essentials edition, Enterprise edition. 
	AuthType string `json:"auth_type,omitempty"`
	// Use the exact http_request string as specified by user, without any automatic insert of headers like Host header. Default value when not specified in API or module is interpreted by ALB Controller as false. 
	ExactHttpRequest bool `json:"exact_http_request,omitempty"`
	// Send an HTTP request to the server. The default GET / HTTP/1.0 may be extended with additional headers or information. For instance, GET /index.htm HTTP/1.1 Host  www.site.com Connection  Close. Default value when not specified in API or module is interpreted by ALB Controller as GET / HTTP/1.0. 
	HttpRequest string `json:"http_request,omitempty"`
	// HTTP request body.
	HttpRequestBody string `json:"http_request_body,omitempty"`
	// Match for a keyword in the first 2Kb of the server header and body response. 
	HttpResponse string `json:"http_response,omitempty"`
	// List of HTTP response codes to match as successful. Default is 2xx. Enum options - HTTP_ANY, HTTP_1XX, HTTP_2XX, HTTP_3XX, HTTP_4XX, HTTP_5XX. Minimum of 1 items required. 
	HttpResponseCode []string `json:"http_response_code"`
	// Match or look for this HTTP response code indicating server maintenance. A successful match results in the server being marked down. Allowed values are 101-599. Maximum of 4 items allowed. 
	MaintenanceCode []int64 `json:"maintenance_code,omitempty"`
	// Match or look for this keyword in the first 2KB of server header and body response indicating server maintenance. A successful match results in the server being marked down. 
	MaintenanceResponse string `json:"maintenance_response,omitempty"`
	// Expected http/https response page size. Allowed values are 2048-16384. 
	ResponseSize int64 `json:"response_size,omitempty"`
	SslAttributes *AlbHealthMonitorSslAttributes `json:"ssl_attributes,omitempty"`
}
