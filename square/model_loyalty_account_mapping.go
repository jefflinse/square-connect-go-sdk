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

// Associates a loyalty account with the buyer's phone number. For more information, see  [Loyalty Overview](/docs/loyalty/overview).
type LoyaltyAccountMapping struct {
	// The Square-assigned ID of the mapping.
	Id string `json:"id,omitempty"`
	// The type of mapping. See [LoyaltyAccountMappingType](#type-loyaltyaccountmappingtype) for possible values
	Type_ string `json:"type"`
	// The phone number, in E.164 format. For example, \"+14155551111\".
	Value string `json:"value"`
	// The timestamp when the mapping was created, in RFC 3339 format.
	CreatedAt string `json:"created_at,omitempty"`
}
