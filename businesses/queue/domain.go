package queue

type Repository interface {
	Publish(name string, data interface{}) error
}