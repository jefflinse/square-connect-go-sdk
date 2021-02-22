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

// Defines parameters in a  [ListSubscriptionEvents](#endpoint-subscriptions-listsubscriptionevents) endpoint request.
type ListSubscriptionEventsRequest struct {
	// A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for the original query.  For more information, see [Pagination](https://developer.squareup.com/docs/docs/working-with-apis/pagination).
	Cursor string `json:"cursor,omitempty"`
	// The upper limit on the number of subscription events to return  in the response.   Default: `200`
	Limit int32 `json:"limit,omitempty"`
}
