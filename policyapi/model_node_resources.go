/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Required node resources to deploy a form factor 
type NodeResources struct {
	// Number of CPU cores required to deploy a form factor. 
	Cpu int64 `json:"cpu,omitempty"`
	// Disk required to deploy a form factor. 
	Disk int64 `json:"disk,omitempty"`
	// Transient storage required to deploy a form factor. 
	EphemeralStorage int64 `json:"ephemeral_storage,omitempty"`
	// Required number of master nodes. 
	NumberOfMasterNodes int64 `json:"number_of_master_nodes,omitempty"`
	// Required number of worker nodes. 
	NumberOfWorkerNodes int64 `json:"number_of_worker_nodes,omitempty"`
	// Memore required to deploy a form factor. 
	Ram int64 `json:"ram,omitempty"`
}
