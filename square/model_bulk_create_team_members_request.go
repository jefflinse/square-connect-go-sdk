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

// Represents a bulk create request for `TeamMember` objects.
type BulkCreateTeamMembersRequest struct {
	// The data which will be used to create the `TeamMember` objects. Each key is the `idempotency_key` that maps to the `CreateTeamMemberRequest`.
	TeamMembers map[string]CreateTeamMemberRequest `json:"team_members"`
}
