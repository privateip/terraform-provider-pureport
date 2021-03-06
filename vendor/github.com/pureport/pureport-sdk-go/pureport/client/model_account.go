/*
 * Pureport Control Plane
 *
 * Pureport API
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package client

type Account struct {
	Demo                      bool              `json:"demo,omitempty"`
	Description               string            `json:"description,omitempty"`
	Href                      string            `json:"href,omitempty"`
	Id                        string            `json:"id,omitempty"`
	Name                      string            `json:"name"`
	Parent                    *Link             `json:"parent,omitempty"`
	ShowChildAccountPricing   bool              `json:"showChildAccountPricing,omitempty"`
	SupportedConnectionGroups []Link            `json:"supportedConnectionGroups,omitempty"`
	Tags                      map[string]string `json:"tags,omitempty"`
}
