/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Schedule to specify day of the week and time to take automated backup
type WeeklyBackupSchedule struct {
	// Schedule type
	ResourceType string `json:"resource_type"`
	// Days of week when backup is taken. 0 - Sunday, 1 - Monday, 2 - Tuesday, 3 - Wednesday ...
	DaysOfWeek []int64 `json:"days_of_week"`
	// Time of day when backup is taken
	HourOfDay int64 `json:"hour_of_day"`
	// Time of day when backup is taken
	MinuteOfDay int64 `json:"minute_of_day"`
}
