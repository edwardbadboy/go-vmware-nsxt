/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type SegmentL2ForwarderSiteSpanInfo struct {
	// Inter-site forwarder status per node.
	InterSiteForwarderStatus []L2ForwarderStatusPerNode `json:"inter_site_forwarder_status,omitempty"`
	// Timestamp when the L2 forwarder remote mac addresses was last updated. 
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	// L2 forwarder remote mac addresses per site for logical switch. 
	RemoteMacsPerSite []L2ForwarderRemoteMacsPerSite `json:"remote_macs_per_site,omitempty"`
	// Policy path of a segment. 
	SegmentPath string `json:"segment_path,omitempty"`
}
