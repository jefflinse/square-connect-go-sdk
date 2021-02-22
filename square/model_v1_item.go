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

// V1Item
type V1Item struct {
	// The item's ID. Must be unique among all entity IDs ever provided on behalf of the merchant. You can never reuse an ID. This value can include alphanumeric characters, dashes (-), and underscores (_).
	Id string `json:"id,omitempty"`
	// The item's name.
	Name string `json:"name,omitempty"`
	// The item's description.
	Description string `json:"description,omitempty"`
	// The item's type. This value is NORMAL for almost all items. See [V1ItemType](#type-v1itemtype) for possible values
	Type_ string `json:"type,omitempty"`
	// The color of the discount's display label in Square Point of Sale, if not the default color. The default color is 9da2a6. See [V1ItemColor](#type-v1itemcolor) for possible values
	Color string `json:"color,omitempty"`
	// The text of the item's display label in Square Point of Sale. Only up to the first five characters of the string are used.
	Abbreviation string `json:"abbreviation,omitempty"`
	// Indicates whether the item is viewable from the merchant's online store (PUBLIC) or PRIVATE. See [V1ItemVisibility](#type-v1itemvisibility) for possible values
	Visibility string `json:"visibility,omitempty"`
	// If true, the item can be added to shipping orders from the merchant's online store.
	AvailableOnline bool `json:"available_online,omitempty"`
	MasterImage *V1ItemImage `json:"master_image,omitempty"`
	Category *V1Category `json:"category,omitempty"`
	// The item's variations. You must specify at least one variation.
	Variations []V1Variation `json:"variations,omitempty"`
	// The modifier lists that apply to the item, if any.
	ModifierLists []V1ModifierList `json:"modifier_lists,omitempty"`
	// The fees that apply to the item, if any.
	Fees []V1Fee `json:"fees,omitempty"`
	// Deprecated. This field is not used.
	Taxable bool `json:"taxable,omitempty"`
	// The ID of the item's category, if any.
	CategoryId string `json:"category_id,omitempty"`
	// If true, the item can be added to pickup orders from the merchant's online store. Default value: false
	AvailableForPickup bool `json:"available_for_pickup,omitempty"`
	// The ID of the CatalogObject in the Connect v2 API. Objects that are shared across multiple locations share the same v2 ID.
	V2Id string `json:"v2_id,omitempty"`
}
