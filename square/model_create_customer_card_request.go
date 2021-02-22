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

// Defines the fields that are included in the request body of a request to the CreateCustomerCard endpoint.
type CreateCustomerCardRequest struct {
	// A card nonce representing the credit card to link to the customer.  Card nonces are generated by the Square Payment Form when customers enter their card information. See [Embedding the payment form](https://developer.squareup.com/docs/payment-form/payment-form-walkthrough) for more information.  __NOTE:__ Card nonces generated by digital wallets (e.g., Apple Pay) cannot be used to create a customer card.
	CardNonce string `json:"card_nonce"`
	BillingAddress *Address `json:"billing_address,omitempty"`
	// The full name printed on the credit card.
	CardholderName string `json:"cardholder_name,omitempty"`
	// An identifying token generated by `SqPaymentForm.verifyBuyer()`. Verification tokens encapsulate customer device information and 3-D Secure challenge results to indicate that Square has verified the buyer identity.
	VerificationToken string `json:"verification_token,omitempty"`
}