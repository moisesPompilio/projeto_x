package input

type CreateUserDTO struct {
	NickName string `json:"nick_name"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
