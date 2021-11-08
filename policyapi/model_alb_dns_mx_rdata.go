/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer DnsMxRdata object
type AlbDnsMxRdata struct {
	// Fully qualified domain name of a mailserver. The host name maps directly to one or more address records in the DNS table, and must not point to any CNAME records (RFC 2181). 
	Host string `json:"host"`
	// The priority field identifies which mail server should be preferred. Allowed values are 0-65535. 
	Priority int64 `json:"priority"`
}
