/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer DnsRrSet object
type AlbDnsRrSet struct {
	Cname *AlbDnsCnameRdata `json:"cname,omitempty"`
	// Fully Qualified Domain Name.
	Fqdn string `json:"fqdn"`
	// IPv6 address in AAAA record.
	Ip6Addresses []AlbDnsAaaaRdata `json:"ip6_addresses,omitempty"`
	// IP address in A record.
	IpAddresses []AlbDnsARdata `json:"ip_addresses,omitempty"`
	// Name Server information in NS record.
	Nses []AlbDnsNsRdata `json:"nses,omitempty"`
	// Time To Live for this DNS record. Allowed values are 0-2147483647. 
	Ttl int64 `json:"ttl"`
	// DNS record type. Enum options - DNS_RECORD_OTHER, DNS_RECORD_A, DNS_RECORD_NS, DNS_RECORD_CNAME, DNS_RECORD_SOA, DNS_RECORD_PTR, DNS_RECORD_HINFO, DNS_RECORD_MX, DNS_RECORD_TXT, DNS_RECORD_RP, DNS_RECORD_DNSKEY, DNS_RECORD_AAAA, DNS_RECORD_SRV, DNS_RECORD_OPT, DNS_RECORD_RRSIG, DNS_RECORD_AXFR, DNS_RECORD_ANY. 
	Type_ string `json:"type"`
}
