/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Provides upgrade summary for a specific site. 
type FederationUpgradeSummary struct {
	// List of component statuses
	ComponentStatus []FederationComponentUpgradeStatus `json:"component_status,omitempty"`
	// This is NSX version for the site.
	CurrentVersion string `json:"current_version,omitempty"`
	// Name of the global manager if present.
	GpmName string `json:"gpm_name,omitempty"`
	// Unique identifier of this resource.
	Id string `json:"id,omitempty"`
	// Indicates the time when the site was upgraded.
	LastUpgradeTimestamp string `json:"last_upgrade_timestamp,omitempty"`
	// Name of the site.
	Name string `json:"name,omitempty"`
	// Status of upgrade
	OverallUpgradeStatus string `json:"overall_upgrade_status,omitempty"`
	// This is the Site Manager generated UUID for every NSX deployment.
	SiteId string `json:"site_id,omitempty"`
	// IP address of the site.
	SiteIp string `json:"site_ip,omitempty"`
	// Type of this site.
	SiteType string `json:"site_type,omitempty"`
	// This is NSX target version for the site, if it is undergoing upgrade.
	TargetVersion string `json:"target_version,omitempty"`
}
