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

// Defines the fields that are included in the response from the [ListSubscriptionEvents](#endpoint-subscriptions-listsubscriptionevents) endpoint.
type ListSubscriptionEventsResponse struct {
	// Information about errors encountered during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// The `SubscriptionEvents` retrieved.
	SubscriptionEvents []SubscriptionEvent `json:"subscription_events,omitempty"`
	// When a response is truncated, it includes a cursor that you can  use in a subsequent request to fetch the next set of events.  If empty, this is the final response.  For more information, see [Pagination](https://developer.squareup.com/docs/docs/working-with-apis/pagination).
	Cursor string `json:"cursor,omitempty"`
}
