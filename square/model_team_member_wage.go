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

// The hourly wage rate that a team member will earn on a `Shift` for doing the job specified by the `title` property of this object.
type TeamMemberWage struct {
	// UUID for this object.
	Id string `json:"id,omitempty"`
	// The `Team Member` that this wage is assigned to.
	TeamMemberId string `json:"team_member_id,omitempty"`
	// The job title that this wage relates to.
	Title string `json:"title,omitempty"`
	HourlyRate *Money `json:"hourly_rate,omitempty"`
}
