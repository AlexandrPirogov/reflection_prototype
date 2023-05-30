package user

type User struct {
	ID      string `json:"-"`
	Login   string `json:"login"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Email   string `json:"email"`
	Pwd     string `json:"pwd"`
}
