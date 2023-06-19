package model

// BASIC STRUCT
type User struct {
	UserID    int    `json:"user_id"`
	FirstName string `json:"f_name"`
	LastName  string `json:"l_name"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

// REQUEST STRUCT
type UserPostReq struct {
	FirstName string `json:"f_name"`
	LastName  string `json:"l_name"`
	Email     string `json:"email"`
}

type UserPutReq struct {
	UserID    int    `json:"user_id"`
	FirstName string `json:"f_name"`
	LastName  string `json:"l_name"`
}

// RESPONS STRUCT
type UserResult struct {
	Status   int    `json:"status" example:"200"`
	Message  string `json:"message" example:"success"`
	Cmd      string `json:"cmd" example:"INSERT/SELECT/UPDATE/DELETE"`
	UserData User   `json:"user_data"`
}

type UserListResult struct {
	Status   int    `json:"status" example:"200"`
	Message  string `json:"message" example:"success"`
	Cmd      string `json:"cmd" example:"INSERT/SELECT/UPDATE/DELETE"`
	UserList []User `json:"user_list"`
}
