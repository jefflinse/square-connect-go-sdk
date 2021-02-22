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

// Contains details on how to fulfill this order.
type OrderFulfillment struct {
	// Unique ID that identifies the fulfillment only within this order.
	Uid string `json:"uid,omitempty"`
	// The type of the fulfillment. See [OrderFulfillmentType](#type-orderfulfillmenttype) for possible values
	Type_ string `json:"type,omitempty"`
	// The state of the fulfillment. See [OrderFulfillmentState](#type-orderfulfillmentstate) for possible values
	State string `json:"state,omitempty"`
	// Application-defined data attached to this fulfillment. Metadata fields are intended to store descriptive references or associations with an entity in another system or store brief information about the object. Square does not process this field; it only stores and returns it in relevant API calls. Do not use metadata to store any sensitive information (personally identifiable information, card details, etc.).  Keys written by applications must be 60 characters or less and must be in the character set `[a-zA-Z0-9_-]`. Entries may also include metadata generated by Square. These keys are prefixed with a namespace, separated from the key with a ':' character.  Values have a max length of 255 characters.  An application may have up to 10 entries per metadata field.  Entries written by applications are private and can only be read or modified by the same application.  See [Metadata](https://developer.squareup.com/docs/build-basics/metadata) for more information.
	Metadata map[string]string `json:"metadata,omitempty"`
	PickupDetails *OrderFulfillmentPickupDetails `json:"pickup_details,omitempty"`
	ShipmentDetails *OrderFulfillmentShipmentDetails `json:"shipment_details,omitempty"`
}
