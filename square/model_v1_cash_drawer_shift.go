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

// Contains details for a single cash drawer shift.
type V1CashDrawerShift struct {
	// The shift's unique ID.
	Id string `json:"id,omitempty"`
	// The shift's current state. See [V1CashDrawerShiftEventType](#type-v1cashdrawershifteventtype) for possible values
	EventType string `json:"event_type,omitempty"`
	// The time when the shift began, in ISO 8601 format.
	OpenedAt string `json:"opened_at,omitempty"`
	// The time when the shift ended, in ISO 8601 format.
	EndedAt string `json:"ended_at,omitempty"`
	// The time when the shift was closed, in ISO 8601 format.
	ClosedAt string `json:"closed_at,omitempty"`
	// The IDs of all employees that were logged into Square Register at some point during the cash drawer shift.
	EmployeeIds []string `json:"employee_ids,omitempty"`
	// The ID of the employee that started the cash drawer shift.
	OpeningEmployeeId string `json:"opening_employee_id,omitempty"`
	// The ID of the employee that ended the cash drawer shift.
	EndingEmployeeId string `json:"ending_employee_id,omitempty"`
	// The ID of the employee that closed the cash drawer shift by auditing the cash drawer's contents.
	ClosingEmployeeId string `json:"closing_employee_id,omitempty"`
	// A description of the cash drawer shift.
	Description string `json:"description,omitempty"`
	StartingCashMoney *V1Money `json:"starting_cash_money,omitempty"`
	CashPaymentMoney *V1Money `json:"cash_payment_money,omitempty"`
	CashRefundsMoney *V1Money `json:"cash_refunds_money,omitempty"`
	CashPaidInMoney *V1Money `json:"cash_paid_in_money,omitempty"`
	CashPaidOutMoney *V1Money `json:"cash_paid_out_money,omitempty"`
	ExpectedCashMoney *V1Money `json:"expected_cash_money,omitempty"`
	ClosedCashMoney *V1Money `json:"closed_cash_money,omitempty"`
	Device *Device `json:"device,omitempty"`
	// All of the events (payments, refunds, and so on) that involved the cash drawer during the shift.
	Events []V1CashDrawerEvent `json:"events,omitempty"`
}
