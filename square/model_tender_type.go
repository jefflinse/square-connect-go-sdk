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
// TenderType : Indicates a tender's type.
type TenderType string

// List of TenderType
const (
	CARD_TenderType TenderType = "CARD"
	CASH_TenderType TenderType = "CASH"
	THIRD_PARTY_CARD_TenderType TenderType = "THIRD_PARTY_CARD"
	SQUARE_GIFT_CARD_TenderType TenderType = "SQUARE_GIFT_CARD"
	NO_SALE_TenderType TenderType = "NO_SALE"
	WALLET_TenderType TenderType = "WALLET"
	OTHER_TenderType TenderType = "OTHER"
)
