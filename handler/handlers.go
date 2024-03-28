package handler

import (
	. "bpsock-go/tags"
)

// Handler is a struct that represents a handler.
type HookHandler struct {
	tag        Tag16
	actionFunc ActionFunc
	cancel     chan string
	data       map[int][]byte
}

// tag
func (h *HookHandler) Tag() Tag16 {
	return h.tag
}

// ActionFunc
func (h *HookHandler) ActionFunc() ActionFunc {
	return h.actionFunc
}

// Cancel
func (h *HookHandler) Cancel() chan string {
	return h.cancel
}

// Data
func (h *HookHandler) Data() map[int][]byte {
	return h.data
}

// HookHandler
func NewHookHandler(tag Tag16, actionFunc ActionFunc) *HookHandler {

	return &HookHandler{
		tag:        tag,
		actionFunc: actionFunc,
		cancel:     make(chan string),
		data:       make(map[int][]byte),
	}
}

//ReqHandler
//TODO: Implement ReqHandler

//ReqPoint
//TODO: Implement ReqPoint
