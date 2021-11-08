/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer GeoLocation object
type AlbGeoLocation struct {
	// Latitude of the location. This is represented as degrees.minutes. The range is from -90.0 (south) to +90.0 (north). Allowed values are -90.0-+90.0. 
	Latitude float32 `json:"latitude,omitempty"`
	// Longitude of the location. This is represented as degrees.minutes. The range is from -180.0 (west) to +180.0 (east). Allowed values are -180.0-+180.0. 
	Longitude float32 `json:"longitude,omitempty"`
	// Location name in the format Country/State/City.
	Name string `json:"name,omitempty"`
	// Location tag string - example  USEast.
	Tag string `json:"tag,omitempty"`
}
