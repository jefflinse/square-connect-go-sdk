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

// Represents one of a business's employees.
type V1Employee struct {
	// The employee's unique ID.
	Id string `json:"id,omitempty"`
	// The employee's first name.
	FirstName string `json:"first_name"`
	// The employee's last name.
	LastName string `json:"last_name"`
	// The ids of the employee's associated roles. Currently, you can specify only one or zero roles per employee.
	RoleIds []string `json:"role_ids,omitempty"`
	// The IDs of the locations the employee is allowed to clock in at.
	AuthorizedLocationIds []string `json:"authorized_location_ids,omitempty"`
	// The employee's email address.
	Email string `json:"email,omitempty"`
	// CWhether the employee is ACTIVE or INACTIVE. Inactive employees cannot sign in to Square Register.Merchants update this field from the Square Dashboard. See [V1EmployeeStatus](#type-v1employeestatus) for possible values
	Status string `json:"status,omitempty"`
	// An ID the merchant can set to associate the employee with an entity in another system.
	ExternalId string `json:"external_id,omitempty"`
	// The time when the employee entity was created, in ISO 8601 format.
	CreatedAt string `json:"created_at,omitempty"`
	// The time when the employee entity was most recently updated, in ISO 8601 format.
	UpdatedAt string `json:"updated_at,omitempty"`
}