/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// This action is performed in HTTP forwarding phase. It is used to inspect the variable of HTTP request, and look up the persistence entry with its value and pool uuid as key. If the persistence entry is found, the HTTP request is forwarded to the recorded backend server according to the persistence entry. If the persistence entry is not found, a new entry is created in the table after backend server is selected. 
type LbVariablePersistenceOnAction struct {
	// The property identifies the load balancer rule action type. 
	Type_ string `json:"type"`
	// If the persistence profile path is not specified, a default persistence table is created per virtual server. Currently, only LBGenericPersistenceProfile is supported. 
	PersistenceProfilePath string `json:"persistence_profile_path,omitempty"`
	// The property is used to enable a hash operation for variable value when composing the persistence key. 
	VariableHashEnabled bool `json:"variable_hash_enabled,omitempty"`
	// The property is the name of variable to be used. It specifies which variable's value of a HTTP Request will be used in the key of persistence entry. The variable can be a built-in variable such as \"_cookie_JSESSIONID\", a customized variable defined in LBVariableAssignmentAction or a captured variable in regular expression such as \"article\". For the full list of built-in variables, please reference the NSX-T Administrator's Guide. 
	VariableName string `json:"variable_name"`
}
