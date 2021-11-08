/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced configuration for Segment
type SegmentAdvancedConfig struct {
	// Connectivity configuration to manually connect (ON) or disconnect (OFF) a Tier1 segment from corresponding Tier1 gateway. Only valid for Tier1 Segments. This property is ignored for L2 VPN extended segments when subnets property is not specified. This property does not apply to VLAN backed segments. VLAN backed segments with connectivity OFF does not affect its layer-2 connectivity. 
	Connectivity string `json:"connectivity,omitempty"`
	// Policy path to IP address pools. 
	AddressPoolPaths []string `json:"address_pool_paths,omitempty"`
	// When set to true, all the ports created on this segment will behave in a hybrid fashion. The hybrid port indicates to NSX that the VM intends to operate in underlay mode, but retains the ability to forward egress traffic to the NSX overlay network. This property is only applicable for segment created with transport zone type OVERLAY_STANDARD. This property cannot be modified after segment is created. 
	Hybrid bool `json:"hybrid,omitempty"`
	// When set to true, any port attached to this logical switch will not be visible through VC/ESX UI 
	InterRouter bool `json:"inter_router,omitempty"`
	// This property is used to enable proximity routing with local egress. When set to true, logical router interface (downlink) connecting Segment to Tier0/Tier1 gateway is configured with prefix-length 32. 
	LocalEgress bool `json:"local_egress,omitempty"`
	// An ordered list of routing policies to forward traffic to the next hop. 
	LocalEgressRoutingPolicies []LocalEgressRoutingEntry `json:"local_egress_routing_policies,omitempty"`
	// Enable multicast on the downlink LRP created to connect the segment to Tier0/Tier1 gateway. 
	Multicast bool `json:"multicast,omitempty"`
	// This profile is applie dto the downlink logical router port created while attaching this semgnet to tier-0 or tier-1. If this field is empty, NDRA profile of the router is applied to the newly created port. 
	NdraProfilePath string `json:"ndra_profile_path,omitempty"`
	// A behaviour required for Firewall As A Service (FaaS) where the segment BUM traffic is confined within the edge node that this segment belongs to. 
	NodeLocalSwitch bool `json:"node_local_switch,omitempty"`
	// ID populated by NSX when NSX on DVPG is used to indicate the source DVPG. Currently, only DVPortgroups are identified as Discovered Segments. The origin_id is the identifier of DVPortgroup from the source vCenter server.
	OriginId string `json:"origin_id,omitempty"`
	// The type of source from where the DVPortgroup is discovered
	OriginType string `json:"origin_type,omitempty"`
	// The name of the switching uplink teaming policy for the Segment. This name corresponds to one of the switching uplink teaming policy names listed in TransportZone associated with the Segment. See transport_zone_path property above for more details. When this property is not specified, the segment will not have a teaming policy associated with it and the host switch's default teaming policy will be used by MP.
	UplinkTeamingPolicyName string `json:"uplink_teaming_policy_name,omitempty"`
	// This URPF mode is applied to the downlink logical router port created while attaching this segment to tier-0 or tier-1. 
	UrpfMode string `json:"urpf_mode,omitempty"`
}
