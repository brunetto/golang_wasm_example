package main

import (
	"fmt"
	"math/rand"
	"syscall/js"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
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

func plot() {
	// from https://github.com/gonum/plot/wiki/Example-plots
	rand.Seed(int64(0))

	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = "Plotutil example"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddLinePoints(p,
		"First", randomPoints(15),
		"Second", randomPoints(15),
		"Third", randomPoints(15))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "points.png"); err != nil {
		panic(err)
	}
}

func randomPoints(n int) plotter.XYs {
	pts := make(plotter.XYs, n)
	for i := range pts {
		if i == 0 {
			pts[i].X = rand.Float64()
		} else {
			pts[i].X = pts[i-1].X + rand.Float64()
		}
		pts[i].Y = pts[i].X + 10*rand.Float64()
	}
	return pts
}
