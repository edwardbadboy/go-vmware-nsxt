/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer SSLKeyAndCertificate object
type AlbsslKeyAndCertificate struct {
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
	// The type of this resource.
	ResourceType string `json:"resource_type,omitempty"`
	// Opaque identifiers meaningful to the API user
	Tags []Tag `json:"tags,omitempty"`
	// Path of its parent
	ParentPath string `json:"parent_path,omitempty"`
	// Absolute path of this object
	Path string `json:"path,omitempty"`
	// This is a UUID generated by the system for realizing the entity object. In most cases this should be same as 'unique_id' of the entity. However, in some cases this can be different because of entities have migrated thier unique identifier to NSX Policy intent objects later in the timeline and did not use unique_id for realization. Realization id is helpful for users to debug data path to correlate the configuration with corresponding intent. 
	RealizationId string `json:"realization_id,omitempty"`
	// Path relative from its parent
	RelativePath string `json:"relative_path,omitempty"`
	// This is a UUID generated by the GM/LM to uniquely identify entites in a federated environment. For entities that are stretched across multiple sites, the same ID will be used on all the stretched sites. 
	UniqueId string `json:"unique_id,omitempty"`
	// subtree for this type within policy tree containing nested elements. 
	Children []ChildPolicyConfigResource `json:"children,omitempty"`
	// Intent objects are not directly deleted from the system when a delete is invoked on them. They are marked for deletion and only when all the realized entities for that intent object gets deleted, the intent object is deleted. Objects that are marked for deletion are not returned in GET call. One can use the search API to get these objects. 
	MarkedForDelete bool `json:"marked_for_delete,omitempty"`
	// Global intent objects cannot be modified by the user. However, certain global intent objects can be overridden locally by use of this property. In such cases, the overridden local values take precedence over the globally defined values for the properties. 
	Overridden bool `json:"overridden,omitempty"`
	// CA certificates in certificate chain.
	CaCerts []AlbCertificateAuthority `json:"ca_certs,omitempty"`
	Certificate *AlbsslCertificate `json:"certificate"`
	// States if the certificate is base64 encoded. Default value when not specified in API or module is interpreted by ALB Controller as false. 
	CertificateBase64 bool `json:"certificate_base64,omitempty"`
	// Creator name.
	CreatedBy string `json:"created_by,omitempty"`
	// Dynamic parameters needed for certificate management profile. 
	DynamicParams []AlbCustomParams `json:"dynamic_params,omitempty"`
	// Enables OCSP Stapling. Allowed in Basic(Allowed values- false) edition, Essentials(Allowed values- false) edition, Enterprise edition. Default value when not specified in API or module is interpreted by ALB Controller as false. 
	EnableOcspStapling bool `json:"enable_ocsp_stapling,omitempty"`
	// Encrypted private key corresponding to the private key (e.g. those generated by an HSM such as Thales nShield). 
	EnckeyBase64 string `json:"enckey_base64,omitempty"`
	// Name of the encrypted private key (e.g. those generated by an HSM such as Thales nShield). 
	EnckeyName string `json:"enckey_name,omitempty"`
	// Format of the Key/Certificate file. Enum options - SSL_PEM, SSL_PKCS12. Default value when not specified in API or module is interpreted by ALB Controller as SSL_PEM. 
	Format string `json:"format,omitempty"`
	// It is a reference to an object of type HardwareSecurityModuleGroup. 
	HardwaresecuritymodulegroupPath string `json:"hardwaresecuritymodulegroup_path,omitempty"`
	// Private key.
	Key string `json:"key,omitempty"`
	// States if the private key is base64 encoded. Default value when not specified in API or module is interpreted by ALB Controller as false. 
	KeyBase64 bool `json:"key_base64,omitempty"`
	KeyParams *AlbsslKeyParams `json:"key_params,omitempty"`
	// Passphrase used to encrypt the private key.
	KeyPassphrase string `json:"key_passphrase,omitempty"`
	// List of labels to be used for granular RBAC. Allowed in Basic edition, Essentials edition, Enterprise edition. 
	Markers []AlbRoleFilterMatchLabel `json:"markers,omitempty"`
	OcspConfig *AlbocspConfig `json:"ocsp_config,omitempty"`
	// Enum options - SSL_CERTIFICATE_FINISHED, SSL_CERTIFICATE_PENDING. Default value when not specified in API or module is interpreted by ALB Controller as SSL_CERTIFICATE_FINISHED. 
	Status string `json:"status,omitempty"`
	// Enum options - SSL_CERTIFICATE_TYPE_VIRTUALSERVICE, SSL_CERTIFICATE_TYPE_SYSTEM, SSL_CERTIFICATE_TYPE_CA. 
	Type_ string `json:"type,omitempty"`
}
