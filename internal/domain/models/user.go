package models

import "time"

type UserRole string

const (
	RoleAdmin       UserRole = "Admin"
	RoleIndividual  UserRole = "Individual"
	RoleLegalEntity UserRole = "LegalEntity"
	RoleScientific  UserRole = "Scientific"
	RoleCorporate   UserRole = "Corporate"
)

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	Role      UserRole  `json:"role"`
	OrgID     *int      `json:"org_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
