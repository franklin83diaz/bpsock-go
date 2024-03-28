package handler

// Handler is an interface that defines the methods that must be implemented by a handler.
type Handler interface {
	Tag() string
	ActionFunc() ActionFunc
	Cancel() chan string
	Data() map[int][]byte
}
