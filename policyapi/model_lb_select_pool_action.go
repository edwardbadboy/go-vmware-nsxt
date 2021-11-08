/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// This action is used to select a pool for matched HTTP request messages. The pool is specified by path. The matched HTTP request messages are forwarded to the specified pool. 
type LbSelectPoolAction struct {
	// The property identifies the load balancer rule action type. 
	Type_ string `json:"type"`
	// Path of load balancer pool.
	PoolId string `json:"pool_id"`
}
