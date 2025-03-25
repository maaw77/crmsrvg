package models

// UserResponse represents the data of a registered user.
type UserResponse struct {
	ID       int    `json:"id,omitempty" db:"id" minimum:"1"`
	Username string `json:"username" db:"username" validate:"required" example:"Some username"`
	Password string `json:"password" db:"password" validate:"required" example:"my_password"`
	Admin    bool   `json:"admin,omitempty" db:"admin" example:"false"`
}

// UserRequest represents the data of a registered user.
type UserRequest struct {
	ID       int    `json:"id,omitempty" db:"id" minimum:"1" swaggerignore:"true"`
	Username string `json:"username" db:"username" validate:"required" example:"Some username"`
	Password string `json:"password" db:"password" validate:"required" example:"my_password"`
	Admin    bool   `json:"admin,omitempty" db:"admin" example:"false" swaggerignore:"true"`
}
