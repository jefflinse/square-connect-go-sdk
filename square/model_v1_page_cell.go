/*
 * Square Connect API
 *
 * Client library for accessing the Square Connect APIs
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package square

// V1PageCell
type V1PageCell struct {
	// The unique identifier of the page the cell is included on.
	PageId string `json:"page_id,omitempty"`
	// The row of the cell. Always an integer between 0 and 4, inclusive.
	Row int32 `json:"row,omitempty"`
	// The column of the cell. Always an integer between 0 and 4, inclusive.
	Column int32 `json:"column,omitempty"`
	// The type of entity represented in the cell (ITEM, DISCOUNT, CATEGORY, or PLACEHOLDER). See [V1PageCellObjectType](#type-v1pagecellobjecttype) for possible values
	ObjectType string `json:"object_type,omitempty"`
	// The unique identifier of the entity represented in the cell. Not present for cells with an object_type of PLACEHOLDER.
	ObjectId string `json:"object_id,omitempty"`
	// For a cell with an object_type of PLACEHOLDER, this value indicates the cell's special behavior. See [V1PageCellPlaceholderType](#type-v1pagecellplaceholdertype) for possible values
	PlaceholderType string `json:"placeholder_type,omitempty"`
}