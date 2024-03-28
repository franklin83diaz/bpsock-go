package handler

// ActionFunc is a function that takes a handler, a Tag type string, and an id type int.
type ActionFunc func(Handler, string, int)
