package models

type UserRole string

const (
	RoleAdmin          UserRole = "Admin"
	RoleIndividual     UserRole = "Individual"
	RoleLegalEntity    UserRole = "LegalEntity"
	RoleScientific     UserRole = "Scientific"
	RoleCorporate      UserRole = "Corporate"
	RoleCorporateAdmin UserRole = "CorporateAdmin"
)

// User представляет пользователя.
type User struct {
	ID       int      `gorm:"primaryKey" json:"id"`
	Name     string   `gorm:"size:255;not null" json:"name"`
	Email    string   `gorm:"unique;size:255;not null" json:"email"`
	Password string   `gorm:"size:255" json:"-"`
	Role     UserRole `gorm:"size:50" json:"role"`
	OrgID    *int     `json:"org_id,omitempty"`
	BaseModel
}

type CreateUserRequest struct {
	Name     string   `json:"name" validate:"required,min=2,max=50"`
	Email    string   `json:"email" validate:"required,email"`
	Password string   `json:"password" validate:"required,min=6"`
	Role     UserRole `json:"role" validate:"required"`
	OrgID    *int     `json:"org_id,omitempty"`
}

type UpdateUserRequest struct {
	Name  *string `json:"name,omitempty" validate:"omitempty,min=2,max=50"`
	Email *string `json:"email,omitempty" validate:"omitempty,email"`
}
