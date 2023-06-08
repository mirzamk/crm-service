package customer

type CustomerDto struct {
	Firstname string `json:"firstname" validate:"required"`
	Lastname  string `json:"lastname" validate:"required"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}
