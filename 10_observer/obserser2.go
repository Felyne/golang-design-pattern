package observer

import "fmt"

type Subject2 struct {
	observers []ObserverFunc
	context   string
}

func NewSubject2() *Subject2 {
	return &Subject2{
		observers: make([]ObserverFunc, 0),
	}
}

func (s *Subject2) Attach(o ObserverFunc) {
	s.observers = append(s.observers, o)
}

func (s *Subject2) notify() {
	for _, o := range s.observers {
		o(s)
	}
}

func (s *Subject2) UpdateContext(context string) {
	s.context = context
	s.notify()
}

type ObserverFunc func(s *Subject2)

type Reader2 struct {
	name string
}

func NewReader2(name string) *Reader2 {
	return &Reader2{
		name: name,
	}
}

func (r *Reader2) Update(s *Subject2) {
	fmt.Printf("%s receive %s\n", r.name, s.context)
}
