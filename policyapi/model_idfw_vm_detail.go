/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Identity Firewall user login/session data for a single VM
type IdfwVmDetail struct {
	LastLoginUserSession *ResourceReference `json:"last_login_user_session,omitempty"`
	// List of user session data.
	UserSessions []IdfwUserSessionData `json:"user_sessions"`
	// Virtual machine (external ID or BIOS UUID) where login/logout event occurred.
	VmExtId string `json:"vm_ext_id"`
	// List of client machine IP addresses.
	VmIpAddresses []string `json:"vm_ip_addresses,omitempty"`
}
