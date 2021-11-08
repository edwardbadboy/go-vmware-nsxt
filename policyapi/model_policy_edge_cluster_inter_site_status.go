/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type PolicyEdgeClusterInterSiteStatus struct {
	// Name of the edge cluster whose status is being reported.
	EdgeClusterName string `json:"edge_cluster_name,omitempty"`
	// Policy path of the edge cluster whose status is being reported.
	EdgeClusterPath string `json:"edge_cluster_path,omitempty"`
	// Timestamp when the edge cluster inter-site status was last updated. 
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// Per edge node inter-site status.
	MemberStatus []PolicyEdgeClusterMemberInterSiteStatus `json:"member_status,omitempty"`
	// Overall status of all edge nodes IBGP status in the edge cluster. 
	OverallStatus string `json:"overall_status,omitempty"`
}
