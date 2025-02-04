/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Cluster restore status
type ClusterRestoreStatus struct {
	// Timestamp when backup was initiated in epoch millisecond
	BackupTimestamp int64 `json:"backup_timestamp,omitempty"`
	// The list of allowed endpoints, based on the current state of the restore process 
	Endpoints []ResourceLink `json:"endpoints,omitempty"`
	// Unique id for backup request
	Id string `json:"id,omitempty"`
	// Instructions for users to reconcile Restore operations
	Instructions []InstructionInfo `json:"instructions,omitempty"`
	// List of actions that are not allowed
	NotAllowedActions []string `json:"not_allowed_actions,omitempty"`
	// Timestamp when restore was completed in epoch millisecond
	RestoreEndTime int64 `json:"restore_end_time,omitempty"`
	// Timestamp when restore was started in epoch millisecond
	RestoreStartTime int64 `json:"restore_start_time,omitempty"`
	Status *GlobalRestoreStatus `json:"status,omitempty"`
	Step *RestoreStep `json:"step,omitempty"`
	// Total number of steps in the entire restore process
	TotalSteps int64 `json:"total_steps,omitempty"`
}
