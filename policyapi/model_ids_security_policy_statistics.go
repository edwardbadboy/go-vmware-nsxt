/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// IDS RSecurity Policy Statistics. 
type IdsSecurityPolicyStatistics struct {
	// Realized id of the section on NSX MP. Policy Manager can create more than one section per SecurityPolicy, in which case this identifier helps to distinguish between the multiple sections created. 
	InternalSectionId string `json:"internal_section_id,omitempty"`
	// Path of the LR on which the section is applied in case of Gateway Firewall. 
	LrPath string `json:"lr_path,omitempty"`
	// Total count for rule statistics
	ResultCount int64 `json:"result_count,omitempty"`
	// List of rule statistics.
	Results []RuleStatistics `json:"results,omitempty"`
}
