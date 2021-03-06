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

// A request for a set of `TeamMemberWage` objects
type ListTeamMemberWagesRequest struct {
	// Filter wages returned to only those that are associated with the specified team member.
	TeamMemberId string `json:"team_member_id,omitempty"`
	// Maximum number of Team Member Wages to return per page. Can range between 1 and 200. The default is the maximum at 200.
	Limit int32 `json:"limit,omitempty"`
	// Pointer to the next page of Employee Wage results to fetch.
	Cursor string `json:"cursor,omitempty"`
}
