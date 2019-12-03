/*
 * Insights Host Inventory REST Interface
 *
 * REST interface for the Insights Platform Host Inventory application.
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package inventory
// TagsOut struct for TagsOut
type TagsOut struct {
	// A number of entries on the current page.
	Count int32 `json:"count,omitempty"`
	// A current page number.
	Page int32 `json:"page,omitempty"`
	// A page size – a number of entries per single page.
	PerPage int32 `json:"per_page,omitempty"`
	// The list of tags on the systems
	Results map[string][]StructuredTag `json:"results,omitempty"`
	// Total number of items in the \"data\" list.
	Total int32 `json:"total,omitempty"`
}
