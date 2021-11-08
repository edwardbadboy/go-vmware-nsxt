/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// NSX global configs for VRNI global collector
type VrniGlobalCollector struct {
	// IP address for the global collector.
	CollectorIp string `json:"collector_ip"`
	// Port for the global collector.
	CollectorPort int32 `json:"collector_port"`
	// Specify the global collector type.
	CollectorType string `json:"collector_type"`
	// Report interval for operation data in seconds.
	ReportInterval int32 `json:"report_interval"`
}
