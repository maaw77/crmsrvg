package models

// User represents the data of a registered user.
type User struct {
	ID       int    `json:"id,omitempty" db:"id" minimum:"1"`
	Username string `json:"username" db:"username" validate:"required" example:"Some username"`
	Password string `json:"password" db:"password" validate:"required" example:"my_password"`
	Admin    bool   `json:"admin,omitempty" db:"admin" example:"true (for admin)"`
}
