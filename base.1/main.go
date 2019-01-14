// +build js,wasm

package main

import (
	"fmt"
	"log"

	"honnef.co/go/js/dom"
)

func main() {
	fmt.Println("Hello, WebAssembly!")
	// mainDiv := dom.GetWindow().Document().GetElementByID("app")
	window := dom.GetWindow()
	if window == nil {
		log.Fatal("window is nil")
	}

	document := window.Document()
	if document == nil {
		log.Fatal("document is nil")
	}

	mainDiv := document.GetElementByID("app")
	if mainDiv == nil {
		log.Fatal("mainDiv is nil")
	}
	// newDiv := dom.GetWindow().Document().CreateElement("div")
	// newDiv.SetTextContent("ciao")
	// mainDiv.AppendChild(newDiv)
}
