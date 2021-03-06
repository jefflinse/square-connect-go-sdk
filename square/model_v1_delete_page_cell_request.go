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

type V1DeletePageCellRequest struct {
	// The row of the cell to clear. Always an integer between 0 and 4, inclusive. Row 0 is the top row.
	Row string `json:"row,omitempty"`
	// The column of the cell to clear. Always an integer between 0 and 4, inclusive. Column 0 is the leftmost column.
	Column string `json:"column,omitempty"`
}
