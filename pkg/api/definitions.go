package api

type NewUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type NewTeamRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type NewTeamMemberRequest struct {
	TeamID int `json:"team_id"`
	UserID int `json:"user_id"`
}

type UpdateManagerRequest struct {
	TeamID int `json:"team_id"`
	UserID int `json:"user_id"`
}

type Team struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Manager     User   `json:"manager"`
	Members     []User `json:"members"`
}

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type GetTeamRequest struct {
	TeamID int `uri:"id" binding:"required"`
}
