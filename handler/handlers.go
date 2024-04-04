package handler

import (
	//lint:ignore ST1001 import tags
	. "bpsock-go/tags"
	"bpsock-go/utils"
)

// Handler is a struct that represents a handler.
type HookHandler struct {
	tag        Tag16
	actionFunc ActionFunc
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
		data:       make(map[int][]byte),
	}
}

// ///////////////////////////////////////////////////////////////////////////
// ReqHandler
type ReqHandler struct {
	tag        Tag16
	actionFunc ActionFunc
	cancel     chan string
	data       map[int][]byte
}

func NewReqHandler(tag Tag8, actionFunc ActionFunc) ReqHandler {

	//Generate a tag ephemera starting with 1
	subTag := utils.AlfanumRandom(7)
	ephemeraTag16 := NewTag16Eph("1" + subTag + tag.Name())

	return ReqHandler{
		tag:        ephemeraTag16,
		actionFunc: actionFunc,
		cancel:     make(chan string),
		data:       make(map[int][]byte),
	}
}

// tag
func (h *ReqHandler) Tag() Tag16 {
	return h.tag

}

// ActionFunc
func (h *ReqHandler) ActionFunc() ActionFunc {
	return h.actionFunc
}

// Data
func (h *ReqHandler) Data() map[int][]byte {
	return h.data
}

// AddData
func (h *ReqHandler) AddData(i int, b []byte) {
	currentData := h.data[i]
	if currentData != nil {
		h.data[i] = append(currentData, b...)
		return
	}
	h.data[i] = b
}

// RemoveData
func (h *ReqHandler) RemoveData(i int) {
	delete(h.data, i)
}

// Cancel
func (h *ReqHandler) Cancel() {
	//TODO: Implement Cancel
	h.cancel <- "cancel"
}

// ///////////////////////////////////////////////////////////////////////////
//ReqPoint
//TODO: Implement ReqPoint
