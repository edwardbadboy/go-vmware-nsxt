/*
 * NSX-T Data Center Policy API
 *
 * VMware NSX-T Data Center Policy REST API
 *
 * API version: 3.2.1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package nsxt

// Advanced load balancer ClientLogStreamingConfig object
type AlbClientLogStreamingConfig struct {
	// IP address or hostnames (FQDNs) of destination servers. If an FQDN is provided, this should be resolvable on Avi Service Engines. Multiple servers are supported by furnishing a comma-separated list of IP addresses or host names, for example, 11.11.11.11,23.12.12.4. Optionally, a separate port can be specified for each external server in the list, for example, 11.11.11.11 234,12.12.12.12 343. 
	ExternalServer string `json:"external_server"`
	// The service port to use for the external servers. If multiple external servers have been specified, the single port number specified here will apply to all those servers for which an explicit port number has not been specified in the external server list. Default value when not specified in API or module is interpreted by ALB Controller as 514. 
	ExternalServerPort int64 `json:"external_server_port,omitempty"`
	FormatConfig *AlbClientLogStreamingFormat `json:"format_config,omitempty"`
	// Type of logs to stream to the external server. Default is LOGS_ALL, i.e., send all logs. Enum options - LOGS_SIGNIFICANT_ONLY, LOGS_UDF_ONLY, LOGS_UDF_SIGNIFICANT, LOGS_ALL. Default value when not specified in API or module is interpreted by ALB Controller as LOGS_ALL. 
	LogTypesToSend string `json:"log_types_to_send,omitempty"`
	// Maximum number of logs per second streamed to the remote server. By default, 100 logs per second are streamed. Set this to zero(0) to not enforce any limit. Default value when not specified in API or module is interpreted by ALB Controller as 100. 
	MaxLogsPerSecond int64 `json:"max_logs_per_second,omitempty"`
	// Protocol to use for streaming logs. Enum options - LOG_STREAMING_PROTOCOL_UDP, LOG_STREAMING_PROTOCOL_SYSLOG_OVER_UDP, LOG_STREAMING_PROTOCOL_TCP, LOG_STREAMING_PROTOCOL_SYSLOG_OVER_TCP, LOG_STREAMING_PROTOCOL_RAW_OVER_UDP, LOG_STREAMING_PROTOCOL_TLS, LOG_STREAMING_PROTOCOL_SYSLOG_OVER_TLS. Default value when not specified in API or module is interpreted by ALB Controller as LOG_STREAMING_PROTOCOL_UDP. 
	Protocol string `json:"protocol,omitempty"`
	SyslogConfig *AlbStreamingSyslogConfig `json:"syslog_config,omitempty"`
}
