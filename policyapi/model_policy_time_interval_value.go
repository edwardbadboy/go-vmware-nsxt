/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Time interval on which firewall schedule will be applicable
type PolicyTimeIntervalValue struct {
	// Time in 24 hour and minutes in multiple of 30. Example, 17:30. 
	EndInterval string `json:"end_interval"`
	// Time in 24 hour and minutes in multiple of 30. Example, 9:00. 
	StartInterval string `json:"start_interval"`
}
