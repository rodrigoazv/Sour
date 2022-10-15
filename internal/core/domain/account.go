package domain

type Account struct {
	ID     string  `json:"id"`
	Amount float64 `json:"amount"`
	Key    string  `json:"key"`
}
