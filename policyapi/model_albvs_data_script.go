/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer VSDataScript object
type AlbvsDataScript struct {
	// Event triggering execution of datascript. Enum options - VS_DATASCRIPT_EVT_HTTP_REQ, VS_DATASCRIPT_EVT_HTTP_RESP, VS_DATASCRIPT_EVT_HTTP_RESP_DATA, VS_DATASCRIPT_EVT_HTTP_LB_FAILED, VS_DATASCRIPT_EVT_HTTP_REQ_DATA, VS_DATASCRIPT_EVT_HTTP_RESP_FAILED, VS_DATASCRIPT_EVT_HTTP_LB_DONE, VS_DATASCRIPT_EVT_HTTP_AUTH, VS_DATASCRIPT_EVT_HTTP_POST_AUTH, VS_DATASCRIPT_EVT_TCP_CLIENT_ACCEPT, VS_DATASCRIPT_EVT_SSL_HANDSHAKE_DONE, VS_DATASCRIPT_EVT_DNS_REQ, VS_DATASCRIPT_EVT_DNS_RESP, VS_DATASCRIPT_EVT_L4_REQUEST, VS_DATASCRIPT_EVT_L4_RESPONSE, VS_DATASCRIPT_EVT_MAX. Allowed in Basic(Allowed values- VS_DATASCRIPT_EVT_HTTP_REQ) edition, Enterprise edition. 
	Evt string `json:"evt"`
	// Datascript to execute when the event triggers.
	Script string `json:"script"`
}
