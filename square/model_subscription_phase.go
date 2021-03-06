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

// Describes a phase in a subscription plan. For more information, see [Set Up and Manage a Subscription Plan](/docs/subscriptions-api/setup-plan).
type SubscriptionPhase struct {
	// The Square-assigned ID of the subscription phase. This field cannot be changed after a `SubscriptionPhase` is created.
	Uid string `json:"uid,omitempty"`
	// The billing cadence of the phase. For example, weekly or monthly. This field cannot be changed after a `SubscriptionPhase` is created. See [SubscriptionCadence](#type-subscriptioncadence) for possible values
	Cadence string `json:"cadence"`
	// The number of `cadence`s the phase lasts. If not set, the phase never ends. Only the last phase can be indefinite. This field cannot be changed after a `SubscriptionPhase` is created.
	Periods int32 `json:"periods,omitempty"`
	RecurringPriceMoney *Money `json:"recurring_price_money"`
	// The position this phase appears in the sequence of phases defined for the plan, indexed from 0. This field cannot be changed after a `SubscriptionPhase` is created.
	Ordinal int64 `json:"ordinal,omitempty"`
}
