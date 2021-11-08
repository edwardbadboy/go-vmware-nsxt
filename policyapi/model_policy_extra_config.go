/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Extra config is intended for supporting vendor specific configuration on the data path, it can be set as key value string pairs on logical switch, logical port or HostSwitch. If it was set on logical switch, it will be inherited automatically by logical ports in it. Also logical port setting will override logical switch setting if specific key was dual set on both logical switch and logical port. 
type PolicyExtraConfig struct {
	ConfigPair *UnboundedKeyValuePair `json:"config_pair"`
}
