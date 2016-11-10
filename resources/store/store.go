package store

type Store interface {
	Connect()
	GetAll(string) (interface{}, error)
	Get(string, string) (interface{}, error)
	Update(string, interface{}) error
	Delete(string) error
	Create(interface{}) (string, error)
}
