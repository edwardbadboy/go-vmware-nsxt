/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// An identity source service that runs OpenLDAP. The service allows selected user accounts defined in OpenLDAP to log into and access NSX-T.
type OpenLdapIdentitySource struct {
	// The server will populate this field when returing the resource. Ignored on PUT and POST.
	Links []ResourceLink `json:"_links,omitempty"`
	// Schema for this resource
	Schema string `json:"_schema,omitempty"`
	Self *SelfResourceLink `json:"_self,omitempty"`
	// The _revision property describes the current revision of the resource. To prevent clients from overwriting each other's changes, PUT operations must include the current _revision of the resource, which clients should obtain by issuing a GET operation. If the _revision provided in a PUT request is missing or stale, the operation will be rejected.
	Revision int32 `json:"_revision,omitempty"`
	// Timestamp of resource creation
	CreateTime int64 `json:"_create_time,omitempty"`
	// ID of the user who created this resource
	CreateUser string `json:"_create_user,omitempty"`
	// Timestamp of last modification
	LastModifiedTime int64 `json:"_last_modified_time,omitempty"`
	// ID of the user who last modified this resource
	LastModifiedUser string `json:"_last_modified_user,omitempty"`
	// Protection status is one of the following: PROTECTED - the client who retrieved the entity is not allowed             to modify it. NOT_PROTECTED - the client who retrieved the entity is allowed                 to modify it REQUIRE_OVERRIDE - the client who retrieved the entity is a super                    user and can modify it, but only when providing                    the request header X-Allow-Overwrite=true. UNKNOWN - the _protection field could not be determined for this           entity. 
	Protection string `json:"_protection,omitempty"`
	// Indicates system owned resource
	SystemOwned bool `json:"_system_owned,omitempty"`
	// Description of this resource
	Description string `json:"description,omitempty"`
	// Defaults to ID if not set
	DisplayName string `json:"display_name,omitempty"`
	// Unique identifier of this resource
	Id string `json:"id,omitempty"`
	ResourceType string `json:"resource_type"`
	// Opaque identifiers meaningful to the API user
	Tags []Tag `json:"tags,omitempty"`
	// After parsing the \"user@domain\", the domain portion is used to select the LDAP identity source to use. Additional domains listed here will also be directed to this LDAP identity source. In Active Directory these are sometimes referred to as Alternative UPN Suffixes.
	AlternativeDomainNames []string `json:"alternative_domain_names,omitempty"`
	// The subtree of the LDAP identity source to search when locating users and groups.
	BaseDn string `json:"base_dn"`
	// The name of the authentication domain. When users log into NSX using an identity of the form \"user@domain\", NSX uses the domain portion to determine which LDAP identity source to use.
	DomainName string `json:"domain_name"`
	// The list of LDAP servers that provide LDAP service for this identity source. Currently, only one LDAP server is supported.
	LdapServers []IdentitySourceLdapServer `json:"ldap_servers,omitempty"`
}
