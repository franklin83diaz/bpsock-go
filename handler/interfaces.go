package handler

import "bpsock-go/tags"

// Handler is an interface that defines the methods that must be implemented by a handler.
type Handler interface {
	Tag() tags.Tag16
	ActionFunc() ActionFunc
	Cancel() chan string
	Data() map[int][]byte
	AddData(i int, b []byte)
	RemoveData(i int)
}
