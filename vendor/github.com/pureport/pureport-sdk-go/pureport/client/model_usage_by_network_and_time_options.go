/*
 * Pureport Control Plane
 *
 * Pureport API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type UsageByNetworkAndTimeOptions struct {
	Date                 *DateFilter `json:"date,omitempty"`
	IncludeChildAccounts bool        `json:"includeChildAccounts,omitempty"`
	TimeUnit             string      `json:"timeUnit,omitempty"`
	TrafficType          string      `json:"trafficType,omitempty"`
}
