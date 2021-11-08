/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced Load Balancer controller node information like node IP and node name.
type AlbControllerNodeInfo struct {
	// Advanced Load Balancer controller node IP configuration is static or DHCP.
	IsDhcp bool `json:"is_dhcp,omitempty"`
	// Advanced Load Balancer controller node IP. 
	NodeIp string `json:"node_ip,omitempty"`
	// Advanced Load Balancer controller node name.
	NodeName string `json:"node_name,omitempty"`
	// Advanced Load Balancer controller node role in cluster.
	NodeRole string `json:"node_role,omitempty"`
	// Advanced Load Balancer controller node start time in its local timezone.
	NodeStartTime int64 `json:"node_start_time,omitempty"`
	// Advanced Load Balancer controller node current state in the cluster.
	NodeState string `json:"node_state,omitempty"`
	// ID of the VM maintained internally. Note: This is automatically generated and cannot be modified. 
	VmId string `json:"vm_id,omitempty"`
}
