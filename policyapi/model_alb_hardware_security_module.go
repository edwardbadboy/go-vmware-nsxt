/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer HardwareSecurityModule object
type AlbHardwareSecurityModule struct {
	Cloudhsm *AlbhsmAwsCloudHsm `json:"cloudhsm,omitempty"`
	// Thales netHSM specific configuration.
	Nethsm []AlbhsmThalesNetHsm `json:"nethsm,omitempty"`
	Rfs *AlbhsmThalesRfs `json:"rfs,omitempty"`
	Sluna *AlbhsmSafenetLuna `json:"sluna,omitempty"`
	// HSM type to use. Enum options - HSM_TYPE_THALES_NETHSM, HSM_TYPE_SAFENET_LUNA, HSM_TYPE_AWS_CLOUDHSM. 
	Type_ string `json:"type"`
}
