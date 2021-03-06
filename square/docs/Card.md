# Card

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Unique ID for this card. Generated by Square. | [optional] [default to null]
**CardBrand** | **string** | The card&#x27;s brand. See [CardBrand](#type-cardbrand) for possible values | [optional] [default to null]
**Last4** | **string** | The last 4 digits of the card number. | [optional] [default to null]
**ExpMonth** | **int64** | The expiration month of the associated card as an integer between 1 and 12. | [optional] [default to null]
**ExpYear** | **int64** | The four-digit year of the card&#x27;s expiration date. | [optional] [default to null]
**CardholderName** | **string** | The name of the cardholder. | [optional] [default to null]
**BillingAddress** | [***Address**](Address.md) |  | [optional] [default to null]
**Fingerprint** | **string** | __Not currently set.__ Intended as a Square-assigned identifier, based on the card number, to identify the card across multiple locations within a single application. | [optional] [default to null]
**CardType** | **string** | The type of the card. The Card object includes this field only in response to Payments API calls. See [CardType](#type-cardtype) for possible values | [optional] [default to null]
**PrepaidType** | **string** | Indicates whether the Card is prepaid or not. The Card object includes this field only in response to Payments API calls. See [CardPrepaidType](#type-cardprepaidtype) for possible values | [optional] [default to null]
**Bin** | **string** | The first six digits of the card number, known as the Bank Identification Number (BIN). Only the Payments API returns this field. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

