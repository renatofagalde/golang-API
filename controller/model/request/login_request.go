package request

type UserLoginRequest struct {
	// o binding é do validator, como está sendo usado através do gin-gonic
	// deve ser usado atraves do binding
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=4,containsany=!@#$%¨&*()"`
}
