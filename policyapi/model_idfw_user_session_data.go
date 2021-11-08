/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Identity Firewall user session data on a client machine (typically a VM). Multiple entries for the same user can be returned if the user logins to multiple sessions on the same VM. 
type IdfwUserSessionData struct {
	// AD Domain of user.
	DomainName string `json:"domain_name"`
	// Identifier of user session data.
	Id string `json:"id,omitempty"`
	// Login time.
	LoginTime int64 `json:"login_time"`
	// Logout time if applicable.  An active user session has no logout time. Non-active user session is stored (up to last 5 most recent entries) per VM and per user. 
	LogoutTime int64 `json:"logout_time,omitempty"`
	// User session source can be one of:   - GI (Guest Introspection)   - ELS (AD Event log server)   - LI (Log Insight) 
	SessionSource string `json:"session_source,omitempty"`
	// AD user ID (may not exist).
	UserId string `json:"user_id,omitempty"`
	// AD user name.
	UserName string `json:"user_name"`
	// User session ID.  This also indicates whether this is VDI / RDSH.
	UserSessionId int32 `json:"user_session_id"`
	// Virtual machine (external ID or BIOS UUID) where login/logout events occurred.
	VmExtId string `json:"vm_ext_id,omitempty"`
}
