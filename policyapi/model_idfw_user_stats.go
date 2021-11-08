/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Identity Firewall user login/session data for a single user.
type IdfwUserStats struct {
	// List of active (still logged in) user login/sessions data (no limit)
	ActiveSessions []IdfwUserSessionData `json:"active_sessions"`
	// Optional list of up to 5 most recent archived (previously logged in) user login/session data. 
	ArchivedSessions []IdfwUserSessionData `json:"archived_sessions,omitempty"`
	// AD user Identifier (String ID)
	UserId string `json:"user_id,omitempty"`
}
