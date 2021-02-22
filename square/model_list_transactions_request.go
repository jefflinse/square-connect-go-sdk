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

// Defines the query parameters that can be included in a request to the [ListTransactions](#endpoint-listtransactions) endpoint.  Deprecated - recommend using [SearchOrders](#endpoint-orders-searchorders)
type ListTransactionsRequest struct {
	// The beginning of the requested reporting period, in RFC 3339 format.  See [Date ranges](#dateranges) for details on date inclusivity/exclusivity.  Default value: The current time minus one year.
	BeginTime string `json:"begin_time,omitempty"`
	// The end of the requested reporting period, in RFC 3339 format.  See [Date ranges](#dateranges) for details on date inclusivity/exclusivity.  Default value: The current time.
	EndTime string `json:"end_time,omitempty"`
	// The order in which results are listed in the response (`ASC` for oldest first, `DESC` for newest first).  Default value: `DESC` See [SortOrder](#type-sortorder) for possible values
	SortOrder string `json:"sort_order,omitempty"`
	// A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for your original query.  See [Paginating results](#paginatingresults) for more information.
	Cursor string `json:"cursor,omitempty"`
}
