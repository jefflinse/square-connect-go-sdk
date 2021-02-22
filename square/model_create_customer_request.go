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

// Defines the body parameters that can be provided in a request to the CreateCustomer endpoint.
type CreateCustomerRequest struct {
	// The idempotency key for the request. See the [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency) guide for more information.
	IdempotencyKey string `json:"idempotency_key,omitempty"`
	// The given (i.e., first) name associated with the customer profile.
	GivenName string `json:"given_name,omitempty"`
	// The family (i.e., last) name associated with the customer profile.
	FamilyName string `json:"family_name,omitempty"`
	// A business name associated with the customer profile.
	CompanyName string `json:"company_name,omitempty"`
	// A nickname for the customer profile.
	Nickname string `json:"nickname,omitempty"`
	// The email address associated with the customer profile.
	EmailAddress string `json:"email_address,omitempty"`
	Address *Address `json:"address,omitempty"`
	// The 11-digit phone number associated with the customer profile.
	PhoneNumber string `json:"phone_number,omitempty"`
	// An optional, second ID used to associate the customer profile with an entity in another system.
	ReferenceId string `json:"reference_id,omitempty"`
	// A custom note associated with the customer profile.
	Note string `json:"note,omitempty"`
	// The birthday associated with the customer profile, in RFC 3339 format. Year is optional, timezone and times are not allowed. For example: `0000-09-01T00:00:00-00:00` indicates a birthday on September 1st. `1998-09-01T00:00:00-00:00` indications a birthday on September 1st __1998__.
	Birthday string `json:"birthday,omitempty"`
}
