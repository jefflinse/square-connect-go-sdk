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

// Describes a slot available for booking, encapsulating appointment segments, the location and starting time.
type Availability struct {
	// The RFC-3339 timestamp specifying the beginning time of the slot available for booking.
	StartAt string `json:"start_at,omitempty"`
	// The ID of the location available for booking.
	LocationId string `json:"location_id,omitempty"`
	// The list of appointment segments available for booking
	AppointmentSegments []AppointmentSegment `json:"appointment_segments,omitempty"`
}
