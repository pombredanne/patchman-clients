/*
 * Insights Host Inventory REST Interface
 *
 * REST interface for the Insights Platform Host Inventory application.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package inventory
// SystemProfileSapSystemOut struct for SystemProfileSapSystemOut
type SystemProfileSapSystemOut struct {
	// The number of items on the current page
	Count int32 `json:"count,omitempty"`
	// The list of sap_system values on the account
	Results []SystemProfileSapSystemOutResults `json:"results,omitempty"`
	// Total number of items
	Total int32 `json:"total,omitempty"`
}
