/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// DHCP option 121 to define classless static route.
type DhcpOption121 struct {
	// Classless static route of DHCP option 121.
	StaticRoutes []ClasslessStaticRoute `json:"static_routes"`
}
