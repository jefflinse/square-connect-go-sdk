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
// ExcludeStrategy : Indicates which products matched by a CatalogPricingRule will be excluded if the pricing rule uses an exclude set.
type ExcludeStrategy string

// List of ExcludeStrategy
const (
	LEAST_EXPENSIVE_ExcludeStrategy ExcludeStrategy = "LEAST_EXPENSIVE"
	MOST_EXPENSIVE_ExcludeStrategy ExcludeStrategy = "MOST_EXPENSIVE"
)