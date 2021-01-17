package main

import "testing"

func TestChain(t *testing.T) {
	c1 := NewProjectManagerChain()
	c2 := NewDepManagerChain()
	c3 := NewGeneralManagerChain()

	c1.SetSuccessor(c2)
	c2.SetSuccessor(c3)

	var c Manager = c1

	c.HandleFeeRequest("bob", 400)
	c.HandleFeeRequest("tom", 1400)
	c.HandleFeeRequest("ada", 10000)
	c.HandleFeeRequest("floar", 400)
	// Output:
	// Project manager permit bob 400 fee request
	// Dep manager permit tom 1400 fee request
	// General manager permit ada 10000 fee request
	// Project manager don't permit floar 400 fee request
}

func TestSensitiveWordFilterChain(t *testing.T) {
	chain := NewSensitiveWordFilterChain()
	chain.AddFilter(NewSexyWordFilter())
	chain.AddFilter(NewAbusiveWordFilter())
	if chain.Filter("sexy") {
		t.Fail()
	}
	if chain.Filter("fuck") {
		t.Fail()
	}
	if false == chain.Filter("good") {
		t.Fail()
	}
}

func TestAdvancedHandlerChain(t *testing.T) {
	chain := NewAdvancedHandlerChain()
	chain.AddHandler(NewHandlerC())
	chain.AddHandler(NewHandlerD())
	chain.Handle()
}
