package main

import "fmt"

// 在 GoF 给出的定义中，如果处理器链上的某个处理器能够处理这个请求，那就不会继续往下传递请求。
// 实际上，职责链模式还有一种变体，那就是请求会被所有的处理器都处理一遍，不存在中途终止的情况。
// 这种变体也有两种实现方式：用链表存储处理器和用数组存储处理器。

type AdvancedHandler interface {
	Handle()
	SetSuccessor(h AdvancedHandler)
}

type implement interface {
	doHandle()
}

type template struct {
	implement
	successor AdvancedHandler
}

func newTemplate(i implement) *template {
	return &template{
		implement: i,
	}
}

func (t *template) Handle() {
	t.implement.doHandle()
	if t.successor != nil {
		t.successor.Handle()
	}
}

func (t *template) SetSuccessor(h AdvancedHandler) {
	t.successor = h
}

type HandlerC struct {
	*template
}

func NewHandlerC() AdvancedHandler {
	h := &HandlerC{}
	template := newTemplate(h)
	h.template = template
	return h
}

func (h *HandlerC) doHandle() {
	fmt.Println("handlerC handle.")
}

type HandlerD struct {
	*template
}

func NewHandlerD() AdvancedHandler {
	h := &HandlerD{}
	template := newTemplate(h)
	h.template = template
	return h
}

func (h *HandlerD) doHandle() {
	fmt.Println("handlerD handle.")
}

type AdvancedHandlerChain struct {
	head AdvancedHandler
	tail AdvancedHandler
}

func NewAdvancedHandlerChain() *AdvancedHandlerChain {
	return &AdvancedHandlerChain{}
}

func (a *AdvancedHandlerChain) AddHandler(h AdvancedHandler) {
	h.SetSuccessor(nil)
	if a.head == nil {
		a.head = h
		a.tail = h
		return
	}
	a.tail.SetSuccessor(h)
	a.tail = h
}

func (a *AdvancedHandlerChain) Handle() {
	if a.head != nil {
		a.head.Handle()
	}
}
