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

type BatchRetrieveCatalogObjectsRequest struct {
	// The IDs of the CatalogObjects to be retrieved.
	ObjectIds []string `json:"object_ids"`
	// If `true`, the response will include additional objects that are related to the requested objects, as follows:  If the `objects` field of the response contains a CatalogItem, its associated CatalogCategory objects, CatalogTax objects, CatalogImage objects and CatalogModifierLists will be returned in the `related_objects` field of the response. If the `objects` field of the response contains a CatalogItemVariation, its parent CatalogItem will be returned in the `related_objects` field of the response.
	IncludeRelatedObjects bool `json:"include_related_objects,omitempty"`
	// The specific version of the catalog objects to be included in the response.  This allows you to retrieve historical versions of objects. The specified version value is matched against the `CatalogObject`s' `version` attribute.
	CatalogVersion int64 `json:"catalog_version,omitempty"`
}
