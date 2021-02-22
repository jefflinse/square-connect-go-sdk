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

// V1EmployeeRole
type V1EmployeeRole struct {
	// The role's unique ID, Can only be set by Square.
	Id string `json:"id,omitempty"`
	// The role's merchant-defined name.
	Name string `json:"name"`
	// The role's permissions. See [V1EmployeeRolePermissions](#type-v1employeerolepermissions) for possible values
	Permissions []string `json:"permissions"`
	// If true, employees with this role have all permissions, regardless of the values indicated in permissions.
	IsOwner bool `json:"is_owner,omitempty"`
	// The time when the employee entity was created, in ISO 8601 format. Is set by Square when the Role is created.
	CreatedAt string `json:"created_at,omitempty"`
	// The time when the employee entity was most recently updated, in ISO 8601 format. Is set by Square when the Role updated.
	UpdatedAt string `json:"updated_at,omitempty"`
}