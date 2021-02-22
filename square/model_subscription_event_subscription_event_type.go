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
// SubscriptionEventSubscriptionEventType : The possible subscription event types.
type SubscriptionEventSubscriptionEventType string

// List of SubscriptionEventSubscriptionEventType
const (
	START_SUBSCRIPTION_SubscriptionEventSubscriptionEventType SubscriptionEventSubscriptionEventType = "START_SUBSCRIPTION"
	PLAN_CHANGE_SubscriptionEventSubscriptionEventType SubscriptionEventSubscriptionEventType = "PLAN_CHANGE"
	STOP_SUBSCRIPTION_SubscriptionEventSubscriptionEventType SubscriptionEventSubscriptionEventType = "STOP_SUBSCRIPTION"
)