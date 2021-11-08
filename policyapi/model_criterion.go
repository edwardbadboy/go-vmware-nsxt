/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Event Criterion is the logical evaluations by which the event may be deemed fulfilled. All the evaluations must be met in order for the criterion to be met (implicit AND). 
type Criterion struct {
	// Criterion Evaluations. 
	Evaluations []Evaluation `json:"evaluations"`
}
