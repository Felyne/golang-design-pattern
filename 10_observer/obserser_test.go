package observer

import "fmt"

func ExampleObserver() {
	subject := NewSubject()
	reader1 := NewReader("reader1")
	reader2 := NewReader("reader2")
	reader3 := NewReader("reader3")
	subject.Attach(reader1)
	subject.Attach(reader2)
	subject.Attach(reader3)

	subject.UpdateContext("observer mode")
	// Output:
	// reader1 receive observer mode
	// reader2 receive observer mode
	// reader3 receive observer mode
}

func ExampleObserverFunc() {
	subject := NewSubject2()
	reader1 := NewReader2("reader1")
	reader2 := NewReader2("reader2")
	reader3 := func(s *Subject2) {
		name := "reader3"
		fmt.Printf("%s receive %s\n", name, s.context)
	}
	subject.Attach(reader1.Update)
	subject.Attach(reader2.Update)
	subject.Attach(reader3)

	subject.UpdateContext("observer mode")
	// Output:
	// reader1 receive observer mode
	// reader2 receive observer mode
	// reader3 receive observer mode
}
