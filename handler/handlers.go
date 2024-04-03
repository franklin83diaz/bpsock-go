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

// AddData
func (h *HookHandler) AddData(i int, b []byte) {
	currentData := h.data[i]
	if currentData != nil {
		h.data[i] = append(currentData, b...)
		return
	}
	h.data[i] = b
}

// RemoveData
func (h *HookHandler) RemoveData(i int) {
	delete(h.data, i)
}

// HookHandler
func NewHookHandler(tag Tag16, actionFunc ActionFunc) HookHandler {

	return HookHandler{
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
