package authDtos

type Creds struct {
	UserName string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type JwtUserResponse struct {
	Id       uint
	Username string
}
