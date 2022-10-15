package ports

type AccountRepository interface {
	Get() error
	Create() error
	Transfer() error
}

type AccountService interface {
	Create() error
	Get() error
	Transfer() error
}
