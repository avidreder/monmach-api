package store

// Store is the interface for data storage, such as a db
type Store interface {
	Connect() error
	GetAll(string, interface{}) error
	GetByKey(string, interface{}, string, interface{}) error
	UpdateByKey(string, map[string]interface{}, string, interface{}) error
	DeleteByKey(string, string, interface{}) error
	Create(string, interface{}) error
}
