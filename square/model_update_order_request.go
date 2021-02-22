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

// Defines the fields that are included in requests to the [UpdateOrder](#endpoint-orders-updateorder) endpoint.
type UpdateOrderRequest struct {
	Order *Order `json:"order,omitempty"`
	// The [dot notation paths](https://developer.squareup.com/docs/orders-api/manage-orders#on-dot-notation) fields to clear. For example, `line_items[uid].note` [Read more about Deleting fields](https://developer.squareup.com/docs/orders-api/manage-orders#delete-fields).
	FieldsToClear []string `json:"fields_to_clear,omitempty"`
	// A value you specify that uniquely identifies this update request  If you're unsure whether a particular update was applied to an order successfully, you can reattempt it with the same idempotency key without worrying about creating duplicate updates to the order. The latest order version will be returned.  See [Idempotency](https://developer.squareup.com/docs/basics/api101/idempotency) for more information.
	IdempotencyKey string `json:"idempotency_key,omitempty"`
}
