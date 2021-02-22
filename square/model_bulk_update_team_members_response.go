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

// Represents a response from a bulk update request, containing the updated `TeamMember` objects or error messages.
type BulkUpdateTeamMembersResponse struct {
	// The successfully updated `TeamMember` objects. Each key is the `team_member_id` that maps to the `UpdateTeamMemberRequest`.
	TeamMembers map[string]UpdateTeamMemberResponse `json:"team_members,omitempty"`
	// The errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}