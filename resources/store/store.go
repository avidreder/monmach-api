package store

// Store is the interface for data storage, such as a db
type Store interface {
	Connect() error
	GetAll(interface{}, string) error
	Get(interface{}) error
	Update(string, int64, map[string]interface{}) error
	Delete(interface{}) error
	Create(interface{}) error
}
