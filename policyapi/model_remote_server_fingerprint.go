/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Remote server
type RemoteServerFingerprint struct {
	// Server port
	Port int64 `json:"port,omitempty"`
	// Remote server hostname or IP address
	Server string `json:"server"`
	// SSH fingerprint of server
	SshFingerprint string `json:"ssh_fingerprint"`
}
