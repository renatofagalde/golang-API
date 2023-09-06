package request

type UserRequest struct {
	// o binding é do validator, como está sendo usado através do gin-gonic
	// deve ser usado atraves do binding
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=4,containsany=!@#$%¨&*()"`
	Name     string `json:"name" binding:"required,min=4,max=50"`
	Age      int8   `json:"age" binding:"required,numeric,min=1,max=100"`
}
