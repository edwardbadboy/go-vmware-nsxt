/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type AntreaTraceflowObservationDropped struct {
	// The type of component. 
	ComponentType string `json:"component_type,omitempty"`
	// UID of the container node that observed a traceflow packet. 
	ContainerNodeId string `json:"container_node_id,omitempty"`
	// The type of observation. AntreaTraceflowObservationDelivered: The packet was delivered to destination Pod properly AntreaTraceflowObservationReceived: The packet was received from another ContainerNode AntreaTraceflowObservationForwarded: The packet was forwarded to next logical node or ContainerNode AntreaTraceflowObservationDropped: The packet was dropped 
	ObservationType string `json:"observation_type"`
	// Timestamp when the observation was collect by Antrea controller. 
	Timestamp int64 `json:"timestamp,omitempty"`
	// The reason of traceflow packet dropped. 
	DropReason string `json:"drop_reason,omitempty"`
	// ID of the hit network policy that makes the traceflow packet dropped. 
	ContainerNetworkPolicyId string `json:"container_network_policy_id,omitempty"`
	// ID of the hit DFW rule that traceflow packet was forwarded. 
	RuleId string `json:"rule_id,omitempty"`
	// The name of component that packets passed through. 
	ComponentName string `json:"component_name,omitempty"`
	// The reject action. Not exist if packet was dropped. 
	RejectAction string `json:"reject_action,omitempty"`
}
