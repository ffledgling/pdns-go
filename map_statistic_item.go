/*
 * PowerDNS Authoritative HTTP API
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 0.0.13
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package pdnsapi

type MapStatisticItem struct {

	// Item name
	Name string `json:"name,omitempty"`

	// Set to \"MapStatisticItem\"
	Type_ string `json:"type,omitempty"`

	// Named values
	Value []SimpleStatisticItem `json:"value,omitempty"`
}
