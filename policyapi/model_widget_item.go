/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Represents a reference to a widget that is held by a container or a multi-widget or a View.
type WidgetItem struct {
	// Aligns widget either left or right.
	Alignment string `json:"alignment,omitempty"`
	Label *Label `json:"label,omitempty"`
	// Represents the vertical span of the widget / container
	Rowspan int32 `json:"rowspan,omitempty"`
	// If true, separates this widget in a container.
	Separator bool `json:"separator,omitempty"`
	// Represents the horizontal span of the widget / container.
	Span int32 `json:"span,omitempty"`
	// Determines placement of widget or container relative to other widgets and containers. The lower the weight, the higher it is in the placement order.
	Weight int32 `json:"weight,omitempty"`
	// Id of the widget configuration that is held by a multi-widget or a container or a view.
	WidgetId string `json:"widget_id"`
}
