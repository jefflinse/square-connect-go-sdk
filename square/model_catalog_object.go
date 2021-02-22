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

// The wrapper object for the Catalog entries of a given object type.  The type of a particular `CatalogObject` is determined by the value of the `type` attribute and only the corresponding data attribute can be set on the `CatalogObject` instance. For example, the following list shows some instances of `CatalogObject` of a given `type` and their corresponding data attribute that can be set: - For a `CatalogObject` of the `ITEM` type, set the `item_data` attribute to yield the `CatalogItem` object. - For a `CatalogObject` of the `ITEM_VARIATION` type, set the `item_variation_data` attribute to yield the `CatalogItemVariation` object. - For a `CatalogObject` of the `MODIFIER` type, set the `modifier_data` attribute to yield the `CatalogModifier` object. - For a `CatalogObject` of the `MODIFIER_LIST` type, set the `modifier_list_data` attribute to yield the `CatalogModifierList` object. - For a `CatalogObject` of the `CATEGORY` type, set the `category_data` attribute to yield the `CatalogCategory` object. - For a `CatalogObject` of the `DISCOUNT` type, set the `discount_data` attribute to yield the `CatalogDiscount` object. - For a `CatalogObject` of the `TAX` type, set the `tax_data` attribute to yield the `CatalogTax` object. - For a `CatalogObject` of the `IMAGE` type, set the `image_data` attribute to yield the `CatalogImageData`  object. - For a `CatalogObject` of the `QUICK_AMOUNTS_SETTINGS` type, set the `quick_amounts_settings_data` attribute to yield the `CatalogQuickAmountsSettings` object. - For a `CatalogObject` of the `PRICING_RULE` type, set the `pricing_rule_data` attribute to yield the `CatalogPricingRule` object. - For a `CatalogObject` of the `TIME_PERIOD` type, set the `time_period_data` attribute to yield the `CatalogTimePeriod` object. - For a `CatalogObject` of the `PRODUCT_SET` type, set the `product_set_data` attribute to yield the `CatalogProductSet`  object. - For a `CatalogObject` of the `SUBSCRIPTION_PLAN` type, set the `subscription_plan_data` attribute to yield the `CatalogSubscriptionPlan` object.   For a more detailed discussion of the Catalog data model, please see the [Design a Catalog](/catalog-api/design-a-catalog) guide.
type CatalogObject struct {
	// The type of this object. Each object type has expected properties expressed in a structured format within its corresponding `*_data` field below. See [CatalogObjectType](#type-catalogobjecttype) for possible values
	Type_ string `json:"type"`
	// An identifier to reference this object in the catalog. When a new `CatalogObject` is inserted, the client should set the id to a temporary identifier starting with a \"`#`\" character. Other objects being inserted or updated within the same request may use this identifier to refer to the new object.  When the server receives the new object, it will supply a unique identifier that replaces the temporary identifier for all future references.
	Id string `json:"id"`
	// Last modification [timestamp](https://developer.squareup.com/docs/build-basics/working-with-dates) in RFC 3339 format, e.g., `\"2016-08-15T23:59:33.123Z\"` would indicate the UTC time (denoted by `Z`) of August 15, 2016 at 23:59:33 and 123 milliseconds.
	UpdatedAt string `json:"updated_at,omitempty"`
	// The version of the object. When updating an object, the version supplied must match the version in the database, otherwise the write will be rejected as conflicting.
	Version int64 `json:"version,omitempty"`
	// If `true`, the object has been deleted from the database. Must be `false` for new objects being inserted. When deleted, the `updated_at` field will equal the deletion time.
	IsDeleted bool `json:"is_deleted,omitempty"`
	// A map (key-value pairs) of application-defined custom attribute values. The value of a key-value pair is a `CatalogCustomAttributeValue` object. The key is the `key` attribute value defined in the associated `CatalogCustomAttributeDefinition` object defined by the application making the request.  If the `CatalogCustomAttributeDefinition` object is defined by another application, the `CatalogCustomAttributeDefinition`'s key attribute value is prefixed by the defining application ID. For example, if the `CatalogCustomAttributeDefinition` has a `key` attribute of `\"cocoa_brand\"` and the defining application ID is `\"abcd1234\"`, the key in the map is `\"abcd1234:cocoa_brand\"` if the application making the request is different from the application defining the custom attribute definition. Otherwise, the key used in the map is simply `\"cocoa_brand\"`.  Application-defined custom attributes that are set at a global (location-independent) level. Custom attribute values are intended to store additional information about a catalog object or associations with an entity in another system. Do not use custom attributes to store any sensitive information (personally identifiable information, card details, etc.).
	CustomAttributeValues map[string]CatalogCustomAttributeValue `json:"custom_attribute_values,omitempty"`
	// The Connect v1 IDs for this object at each location where it is present, where they differ from the object's Connect V2 ID. The field will only be present for objects that have been created or modified by legacy APIs.
	CatalogV1Ids []CatalogV1Id `json:"catalog_v1_ids,omitempty"`
	// If `true`, this object is present at all locations (including future locations), except where specified in the `absent_at_location_ids` field. If `false`, this object is not present at any locations (including future locations), except where specified in the `present_at_location_ids` field. If not specified, defaults to `true`.
	PresentAtAllLocations bool `json:"present_at_all_locations,omitempty"`
	// A list of locations where the object is present, even if `present_at_all_locations` is `false`.
	PresentAtLocationIds []string `json:"present_at_location_ids,omitempty"`
	// A list of locations where the object is not present, even if `present_at_all_locations` is `true`.
	AbsentAtLocationIds []string `json:"absent_at_location_ids,omitempty"`
	// Identifies the `CatalogImage` attached to this `CatalogObject`.
	ImageId string `json:"image_id,omitempty"`
	ItemData *CatalogItem `json:"item_data,omitempty"`
	CategoryData *CatalogCategory `json:"category_data,omitempty"`
	ItemVariationData *CatalogItemVariation `json:"item_variation_data,omitempty"`
	TaxData *CatalogTax `json:"tax_data,omitempty"`
	DiscountData *CatalogDiscount `json:"discount_data,omitempty"`
	ModifierListData *CatalogModifierList `json:"modifier_list_data,omitempty"`
	ModifierData *CatalogModifier `json:"modifier_data,omitempty"`
	TimePeriodData *CatalogTimePeriod `json:"time_period_data,omitempty"`
	ProductSetData *CatalogProductSet `json:"product_set_data,omitempty"`
	PricingRuleData *CatalogPricingRule `json:"pricing_rule_data,omitempty"`
	ImageData *CatalogImage `json:"image_data,omitempty"`
	MeasurementUnitData *CatalogMeasurementUnit `json:"measurement_unit_data,omitempty"`
	SubscriptionPlanData *CatalogSubscriptionPlan `json:"subscription_plan_data,omitempty"`
	ItemOptionData *CatalogItemOption `json:"item_option_data,omitempty"`
	ItemOptionValueData *CatalogItemOptionValue `json:"item_option_value_data,omitempty"`
	CustomAttributeDefinitionData *CatalogCustomAttributeDefinition `json:"custom_attribute_definition_data,omitempty"`
	QuickAmountsSettingsData *CatalogQuickAmountsSettings `json:"quick_amounts_settings_data,omitempty"`
}
