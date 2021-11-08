/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type PolicyTepCsvListResult struct {
	// File name set by HTTP server if API  returns CSV result as a file.
	FileName string `json:"file_name,omitempty"`
	// Timestamp when the data was last updated; unset if data source has never updated the data.
	LastUpdateTimestamp int64 `json:"last_update_timestamp,omitempty"`
	Results []PolicyTepTableCsvRecord `json:"results,omitempty"`
}
