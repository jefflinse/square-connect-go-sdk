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

// A request to calculate the points that a buyer can earn from  a specified purchase.
type CalculateLoyaltyPointsRequest struct {
	// The `order` ID for which to calculate the points. Specify this field if your application uses the Orders API to process orders. Otherwise, specify the `transaction_amount`.
	OrderId string `json:"order_id,omitempty"`
	TransactionAmountMoney *Money `json:"transaction_amount_money,omitempty"`
}
