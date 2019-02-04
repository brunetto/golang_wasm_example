package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	fmt.Println("Hello, WebAssembly!")
	r := root{
		el: js.Global().Get("document").Call("getElementById", "root-div"),
	}

	fmt.Println("Created root div")

	d := newDiv("pippo")
	d.appendTo(r.el)
}

type Element struct {
	tag    string
	params map[string]string
}

func (el *Element) create() js.Value {
	e := js.Global().Get("document").Call("createElement", el.tag)
	for attr, value := range el.params {
		e.Set(attr, value)
	}

	fmt.Println("div created")
	return e
}

type root struct {
	el js.Value
}

type div struct {
	el   js.Value
	text string
}

func newDiv(text string) *div {
	el := Element{
		tag:    "div",
		params: map[string]string{},
	}
	d := el.create()
	d.Set("innerText", text)
	fmt.Println("text added")
	return &div{el: d, text: text}
}

func (d *div) appendTo(parent js.Value) {
	parent.Call("appendChild", d.el)
	fmt.Println("div appended")
}
