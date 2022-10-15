package domain

type User struct {
	ID       string  `json:"id"`
	Nickname string  `json:"name"`
	Cpf      string  `json:"state"`
	Password string  `json:"password"`
	Account  Account `json:"account"`
}
