package actor

type ActorDto struct {
	Username   string `json:"username"`
	Role       string `json:"role"`
	IsVerified string `json:"is_verified"`
	IsActive   string `json:"is_active"`
}

type ApprovalDto struct {
	ID     uint     `json:"approval_id"`
	Admin  ActorDto `json:"admin"`
	Status string   `json:"status"`
}
