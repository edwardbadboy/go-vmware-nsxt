/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// The capacity contains basic information and load balancer entity usages and capacity for the given edge node. 
type LbEdgeNodeUsage struct {
	// The property identifies the node path for load balancer node usage. For example, node_path=/infra/sites/default/enforcement-points/default /edge-clusters/85175e0b-4d74-461d-83e1-f3b785adef9c/edge-nodes /86e077c0-449f-11e9-87c8-02004eb37029. 
	NodePath string `json:"node_path"`
	// The property identifies the load balancer node usage type. 
	ResourceType string `json:"resource_type"`
	// The count of large load balancer services configured on the node. 
	CurrentLargeLoadBalancerCount int64 `json:"current_large_load_balancer_count,omitempty"`
	// The current load balancer credits means the current credits used on the node. For example, configuring a medium load balancer on a node consumes 10 credits. If there are 2 medium instances configured on a node, the current credit number is 2 * 10 = 20. 
	CurrentLoadBalancerCredits int64 `json:"current_load_balancer_credits,omitempty"`
	// The count of medium load balancer services configured on the node. 
	CurrentMediumLoadBalancerCount int64 `json:"current_medium_load_balancer_count,omitempty"`
	// The count of pools configured on the node. 
	CurrentPoolCount int64 `json:"current_pool_count,omitempty"`
	// The count of pool members configured on the node. 
	CurrentPoolMemberCount int64 `json:"current_pool_member_count,omitempty"`
	// The count of small load balancer services configured on the node. 
	CurrentSmallLoadBalancerCount int64 `json:"current_small_load_balancer_count,omitempty"`
	// The count of virtual servers configured on the node. 
	CurrentVirtualServerCount int64 `json:"current_virtual_server_count,omitempty"`
	// The count of xlarge load balancer services configured on the node. 
	CurrentXlargeLoadBalancerCount int64 `json:"current_xlarge_load_balancer_count,omitempty"`
	// The path of edge cluster which contains the edge node. 
	EdgeClusterPath string `json:"edge_cluster_path,omitempty"`
	// The form factor of the given edge node. 
	FormFactor string `json:"form_factor,omitempty"`
	// The load balancer credit capacity means the maximum credits which can be used for load balancer configuration for the given edge node. 
	LoadBalancerCreditCapacity int64 `json:"load_balancer_credit_capacity,omitempty"`
	// Pool member capacity means maximum number of pool members which can be configured on the given edge node. 
	PoolMemberCapacity int64 `json:"pool_member_capacity,omitempty"`
	// The remaining count of large load balancer services which can be configured on the given edge node. 
	RemainingLargeLoadBalancerCount int64 `json:"remaining_large_load_balancer_count,omitempty"`
	// The remaining count of medium load balancer services which can be configured on the given edge node. 
	RemainingMediumLoadBalancerCount int64 `json:"remaining_medium_load_balancer_count,omitempty"`
	// The remaining count of small load balancer services which can be configured on the given edge node. 
	RemainingSmallLoadBalancerCount int64 `json:"remaining_small_load_balancer_count,omitempty"`
	// The remaining count of xlarge load balancer services which can be configured on the given edge node. 
	RemainingXlargeLoadBalancerCount int64 `json:"remaining_xlarge_load_balancer_count,omitempty"`
	// The severity calculation is based on current credit usage percentage of load balancer for one node. 
	Severity string `json:"severity,omitempty"`
	// The usage percentage of the edge node for load balancer. The value is the larger value between load balancer credit usage percentage and pool member usage percentage for the edge node. 
	UsagePercentage float32 `json:"usage_percentage,omitempty"`
}
