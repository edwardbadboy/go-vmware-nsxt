/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer DnsQueryNameMatch object
type AlbDnsQueryNameMatch struct {
	// Criterion to use for string matching the DNS query domain name in the question section. Enum options - BEGINS_WITH, DOES_NOT_BEGIN_WITH, CONTAINS, DOES_NOT_CONTAIN, ENDS_WITH, DOES_NOT_END_WITH, EQUALS, DOES_NOT_EQUAL, REGEX_MATCH, REGEX_DOES_NOT_MATCH. Allowed in Basic(Allowed values- BEGINS_WITH,DOES_NOT_BEGIN_WITH,CONTAINS,DOES_NOT_CONTAIN,ENDS_WITH,DOES_NOT_END_WITH,EQUALS,DOES_NOT_EQUAL) edition, Essentials(Allowed values- BEGINS_WITH,DOES_NOT_BEGIN_WITH,CONTAINS,DOES_NOT_CONTAIN,ENDS_WITH,DOES_NOT_END_WITH,EQUALS,DOES_NOT_EQUAL) edition, Enterprise edition. 
	MatchCriteria string `json:"match_criteria"`
	// Domain name to match against that specified in the question section of the DNS query. 
	QueryDomainNames []string `json:"query_domain_names,omitempty"`
	// path of the string group(s) for matching against DNS query domain name in the question section. It is a reference to an object of type StringGroup. 
	StringGroupPaths []string `json:"string_group_paths,omitempty"`
}
