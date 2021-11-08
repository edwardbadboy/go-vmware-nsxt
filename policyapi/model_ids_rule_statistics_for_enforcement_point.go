/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// IDS Rule statistics for a specfic enforcement point.
type IdsRuleStatisticsForEnforcementPoint struct {
	// IDS Rule statistics for a single enforcement point
	EnforcementPoint string `json:"enforcement_point,omitempty"`
	Statistics *IdsRuleStatistics `json:"statistics,omitempty"`
}
