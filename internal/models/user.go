package models

// User represents the data of a registered user.
type User struct {
	// ID of the database entry
	//
	// required:false
	// min:1
	ID int `json:"id,omitempty" db:"id"`

	// Username of a a registered user
	//
	// required: true
	// example: Some username
	Username string `json:"username" db:"username"`

	// Password of a a registered user
	//
	// required: true
	// example: my_password
	Password string `json:"password" db:"password"`

	// User's status
	// required: false
	// example: true (for admin)
	Admin bool `json:"admin,omitempty" db:"admin"`
}
