package store

// Store is the interface for data storage, such as a db
type Store interface {
	Connect()
	GetAll(string) (interface{}, error)
	Get(string, string) (interface{}, error)
	Update(string, interface{}) error
	Delete(string) error
	Create(string, map[string]interface{}) (string, error)
}
