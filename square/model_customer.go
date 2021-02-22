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

// Represents a Square customer profile, which can have one or more cards on file associated with it.
type Customer struct {
	// A unique Square-assigned ID for the customer profile.
	Id string `json:"id,omitempty"`
	// The timestamp when the customer profile was created, in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`
	// The timestamp when the customer profile was last updated, in RFC 3339 format.
	UpdatedAt string `json:"updated_at,omitempty"`
	// Payment details of cards stored on file for the customer profile.
	Cards []Card `json:"cards,omitempty"`
	// The given (i.e., first) name associated with the customer profile.
	GivenName string `json:"given_name,omitempty"`
	// The family (i.e., last) name associated with the customer profile.
	FamilyName string `json:"family_name,omitempty"`
	// A nickname for the customer profile.
	Nickname string `json:"nickname,omitempty"`
	// A business name associated with the customer profile.
	CompanyName string `json:"company_name,omitempty"`
	// The email address associated with the customer profile.
	EmailAddress string `json:"email_address,omitempty"`
	Address *Address `json:"address,omitempty"`
	// The 11-digit phone number associated with the customer profile.
	PhoneNumber string `json:"phone_number,omitempty"`
	// The birthday associated with the customer profile, in RFC 3339 format. Year is optional, timezone and times are not allowed. For example: `0000-09-01T00:00:00-00:00` indicates a birthday on September 1st. `1998-09-01T00:00:00-00:00` indications a birthday on September 1st __1998__.
	Birthday string `json:"birthday,omitempty"`
	// An optional, second ID used to associate the customer profile with an entity in another system.
	ReferenceId string `json:"reference_id,omitempty"`
	// A custom note associated with the customer profile.
	Note string `json:"note,omitempty"`
	Preferences *CustomerPreferences `json:"preferences,omitempty"`
	// The customer groups and segments the customer belongs to. This deprecated field has been replaced with  the dedicated `group_ids` for customer groups and the dedicated `segment_ids` field for customer segments. You can retrieve information about a given customer group and segment respectively using the Customer Groups API and Customer Segments API.
	Groups []CustomerGroupInfo `json:"groups,omitempty"`
	// A creation source represents the method used to create the customer profile. See [CustomerCreationSource](#type-customercreationsource) for possible values
	CreationSource string `json:"creation_source,omitempty"`
	// The IDs of customer groups the customer belongs to.
	GroupIds []string `json:"group_ids,omitempty"`
	// The IDs of segments the customer belongs to.
	SegmentIds []string `json:"segment_ids,omitempty"`
}
