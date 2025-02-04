/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Tier-0 vrf configuration.
type Tier0VrfConfig struct {
	EvpnL2VniConfig *VrfEvpnL2VniConfig `json:"evpn_l2_vni_config,omitempty"`
	// L3 VNI associated with the VRF for overlay traffic of ethernet virtual private network (EVPN). It must be unique and available from the VNI pool defined for EVPN service. It is required for VRF to participate in the EVPN service in INLINE mode. 
	EvpnTransitVni int32 `json:"evpn_transit_vni,omitempty"`
	// Route distinguisher with format in IPAddress:<number> or ASN:<number>.
	RouteDistinguisher string `json:"route_distinguisher,omitempty"`
	// Route targets.
	RouteTargets []VrfRouteTargets `json:"route_targets,omitempty"`
	// Default tier0 path. Cannot be modified after realization. 
	Tier0Path string `json:"tier0_path"`
}
