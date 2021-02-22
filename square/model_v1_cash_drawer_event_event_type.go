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

type V1CashDrawerEventEventType string

// List of V1CashDrawerEventEventType
const (
	NO_SALE_V1CashDrawerEventEventType V1CashDrawerEventEventType = "NO_SALE"
	CASH_TENDER_PAYMENT_V1CashDrawerEventEventType V1CashDrawerEventEventType = "CASH_TENDER_PAYMENT"
	OTHER_TENDER_PAYMENT_V1CashDrawerEventEventType V1CashDrawerEventEventType = "OTHER_TENDER_PAYMENT"
	CASH_TENDER_CANCELED_PAYMENT_V1CashDrawerEventEventType V1CashDrawerEventEventType = "CASH_TENDER_CANCELED_PAYMENT"
	OTHER_TENDER_CANCELED_PAYMENT_V1CashDrawerEventEventType V1CashDrawerEventEventType = "OTHER_TENDER_CANCELED_PAYMENT"
	CASH_TENDER_REFUND_V1CashDrawerEventEventType V1CashDrawerEventEventType = "CASH_TENDER_REFUND"
	OTHER_TENDER_REFUND_V1CashDrawerEventEventType V1CashDrawerEventEventType = "OTHER_TENDER_REFUND"
	PAID_IN_V1CashDrawerEventEventType V1CashDrawerEventEventType = "PAID_IN"
	PAID_OUT_V1CashDrawerEventEventType V1CashDrawerEventEventType = "PAID_OUT"
)
