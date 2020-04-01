package wtk

import (
	"fmt"
	"github.com/worldiety/wtk/dom"
	"log"
	"runtime"
	"strings"
)

var Root = newWindow()

type Window struct {
	window dom.Window
}

func (w Window) attach(parent View) {
}

func (w Window) detach() {
}

func (w Window) parent() View {
	return nil
}

func (w Window) node() dom.Element {
	return w.window.Document().Body()
}

func (w Window) Release() {
}

func newWindow() Window {
	return Window{window: dom.GetWindow()}
}

func (w Window) RemoveAll() {
	w.window.Document().Body().SetInnerHTML("")
}

func (w Window) AddView(v View) {
	v.attach(w)
	w.node().AppendChild(v.node())
}

func Run(target View, init func()) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("GOT THE PANIC")
			b := make([]byte, 2048) // adjust buffer size to be larger than expected stack
			n := runtime.Stack(b, false)
			s := fmt.Sprintf("%v:\n", r) + string(b[:n])
			target.node().SetTextContent("")
			lines := strings.Split(s, "\n")
			for _, line := range lines {
				e := dom.CreateElement("p")
				//	e.Style().AddClass("stacktraceLine")
				e.SetTextContent(line)
				target.node().AppendChild(e)
			}

		}
	}()
	init()
}
