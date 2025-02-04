/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Dscp bit config
type DscpBit struct {
	// The method for indicating the existence of INT header.
	IndicatorType string `json:"indicator_type"`
	// A DSCP bit is allocated to indicate the existence of INT header. It takes effect only when the INT indicator mode is DSCP_BIT. The user should guarantee that the given DSCP bit is specifically allocated for INT. 
	DscpBit int32 `json:"dscp_bit"`
}
