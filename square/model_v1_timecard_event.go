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

// V1TimecardEvent
type V1TimecardEvent struct {
	// The event's unique ID.
	Id string `json:"id,omitempty"`
	// The ID of the timecard to list events for. See [V1TimecardEventEventType](#type-v1timecardeventeventtype) for possible values
	EventType string `json:"event_type,omitempty"`
	// The time the employee clocked in, in ISO 8601 format.
	ClockinTime string `json:"clockin_time,omitempty"`
	// The time the employee clocked out, in ISO 8601 format.
	ClockoutTime string `json:"clockout_time,omitempty"`
	// The time when the event was created, in ISO 8601 format.
	CreatedAt string `json:"created_at,omitempty"`
}
