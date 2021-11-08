/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer LdapUserBindSettings object
type AlbLdapUserBindSettings struct {
	// LDAP user DN pattern is used to bind LDAP user after replacing the user token with real username. 
	DnTemplate string `json:"dn_template,omitempty"`
	// LDAP token is replaced with real user name in the user DN pattern. Default value when not specified in API or module is interpreted by ALB Controller as <user>. 
	Token string `json:"token,omitempty"`
	// LDAP user attributes to fetch on a successful user bind.
	UserAttributes []string `json:"user_attributes,omitempty"`
	// LDAP user id attribute is the login attribute that uniquely identifies a single user record. 
	UserIdAttribute string `json:"user_id_attribute,omitempty"`
}
