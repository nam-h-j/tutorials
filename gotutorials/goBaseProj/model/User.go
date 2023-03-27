package model

import "time"

// BASIC STRUCT
type User struct {
	ID        int       `json:"id"`
	FirstName string    `json:"f_name"`
	LastName  string    `json:"l_name"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

// RESPONS STRUCT
type UserResult struct {
	Status   int    `json:"status" example:"200"`
	Message  string `json:"message" example:"success"`
	Cmd      string `json:"cmd" example:"INSERT/SELECT/UPDATE/DELETE"`
	UserData User   `json:"user_data"`
}
