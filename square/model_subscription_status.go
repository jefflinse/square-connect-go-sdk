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
// SubscriptionStatus : Possible subscription status values.
type SubscriptionStatus string

// List of SubscriptionStatus
const (
	DEFAULT_SUBSCRIPTION_STATUS_DO_NOT_USE_SubscriptionStatus SubscriptionStatus = "DEFAULT_SUBSCRIPTION_STATUS_DO_NOT_USE"
	PENDING_SubscriptionStatus SubscriptionStatus = "PENDING"
	ACTIVE_SubscriptionStatus SubscriptionStatus = "ACTIVE"
	CANCELED_SubscriptionStatus SubscriptionStatus = "CANCELED"
)
