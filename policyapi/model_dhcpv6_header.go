/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

type Dhcpv6Header struct {
	// This is used to specify the DHCP v6 message. To request the assignment of one or more IPv6 addresses, a client first locates a DHCP server and then requests the assignment of addresses and other configuration information from the server. The client sends a Solicit message to the All_DHCP_Relay_Agents_and_Servers address to find available DHCP servers. Any server that can meet the client's requirements responds with an Advertise message. The client then chooses one of the servers and sends a Request message to the server asking for confirmed assignment of addresses and other configuration information. The server responds with a Reply message that contains the confirmed addresses and configuration. SOLICIT - A client sends a Solicit message to locate servers. ADVERTISE - A server sends and Advertise message to indicate that it is available. REQUEST - A client sends a Request message to request configuration parameters. REPLY - A server sends a Reply message containing assigned addresses and configuration parameters.
	MsgType string `json:"msg_type,omitempty"`
}
