package main

import "fmt"

type VisitorFunc func(*Info, error) error

type Visitor interface {
	Visit(VisitorFunc) error
}

type Info struct {
	Namespace   string
	Name        string
	OtherThings string
}

func (info *Info) Visit(fn VisitorFunc) error {
	return fn(info, nil)
}

func NameVisitor(info *Info, err error) error {
	fmt.Println("NameVisitor() before call function")
	if err == nil {
		fmt.Printf("==> Name=%s, NameSpace=%s\n", info.Name, info.Namespace)
	}
	fmt.Println("NameVisitor() after call function")
	return err
}

func OtherThingsVisitor(info *Info, err error) error {
	fmt.Println("OtherThingsVisitor() before call function")
	if err == nil {
		fmt.Printf("==> OtherThings=%s\n", info.OtherThings)
	}
	fmt.Println("OtherThingsVisitor() after call function")
	return err
}

func LoadFile(info *Info, err error) error {
	info.Name = "Hao Chen"
	info.Namespace = "MegaEase"
	info.OtherThings = "We are running as remote team."
	return nil
}

type DecoratedVisitor struct {
	visitor    Visitor
	decorators []VisitorFunc
}

func NewDecoratedVisitor(v Visitor, fn ...VisitorFunc) Visitor {
	if len(fn) == 0 {
		return v
	}
	return DecoratedVisitor{v, fn}
}

// Visit implements Visitor
func (v DecoratedVisitor) Visit(fn VisitorFunc) error {
	return v.visitor.Visit(func(info *Info, err error) error {
		if err != nil {
			return err
		}
		if err := fn(info, nil); err != nil {
			return err
		}
		for i := range v.decorators {
			if err := v.decorators[i](info, nil); err != nil {
				return err
			}
		}
		return nil
	})
}

func main() {
	info := Info{}
	var v Visitor = &info
	v = NewDecoratedVisitor(v, NameVisitor, OtherThingsVisitor)

	v.Visit(LoadFile)
}
