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
	data       map[int][]byte
}

func NewReqHandler(tag Tag8, actionFunc ActionFunc) ReqHandler {

	//Generate a tag ephemera
	subTag := utils.AlfanumRandom(7)
	ephemeraTag16 := NewTag16Eph(subTag + tag.Name())

	return ReqHandler{
		tag:        ephemeraTag16,
		actionFunc: actionFunc,

		data: make(map[int][]byte),
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

// ///////////////////////////////////////////////////////////////////////////
// ReqPoint
type ReqPoint struct {
	tag        Tag16
	actionFunc ActionFunc
	Cancel     chan string
	data       map[int][]byte
}

func NewReqPoint(tag Tag16, actionFunc ActionFunc) ReqPoint {
	return ReqPoint{
		tag:        tag,
		actionFunc: actionFunc,
		Cancel:     make(chan string),
		data:       make(map[int][]byte),
	}
}

// tag
func (h *ReqPoint) Tag() Tag16 {
	return h.tag
}

// ActionFunc
func (h *ReqPoint) ActionFunc() ActionFunc {
	return h.actionFunc
}

// Data
func (h *ReqPoint) Data() map[int][]byte {
	return h.data
}

// AddData
func (h *ReqPoint) AddData(i int, b []byte) {
	currentData := h.data[i]
	if currentData != nil {
		h.data[i] = append(currentData, b...)
		return
	}
	h.data[i] = b
}

// RemoveData
func (h *ReqPoint) RemoveData(i int) {
	delete(h.data, i)
}

// ///////////////////////////////////////////////////////////////////////////
