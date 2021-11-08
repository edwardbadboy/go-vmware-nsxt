/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Single route redistribution rule
type Tier0RouteRedistributionRule struct {
	// Each rule can have more than one destinations. If destinations not specified for a given rule, default destionation will be BGP 
	Destinations []string `json:"destinations,omitempty"`
	// Rule name
	Name string `json:"name,omitempty"`
	// Route map to be associated with the redistribution rule
	RouteMapPath string `json:"route_map_path,omitempty"`
	// List of redistribution types
	RouteRedistributionTypes []string `json:"route_redistribution_types"`
}
