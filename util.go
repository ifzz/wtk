package wtk

import (
	"strconv"
	"strings"
	"syscall/js"
)

var htmlId = 1

// assertNotAttached bails out if parent is not nil
func assertNotAttached(v View) {
	if v.parent() != nil {
		panic("invalid state: view is already attached")
	}
}

// assertAttached bails out if parent is nil
func assertAttached(v View) {
	if v.parent() == nil {
		panic("invalid state: view is not attached")
	}
}

func floatToPx(v float64) string {
	return strconv.Itoa(int(v)) + "px"
}

func nextId() string {
	htmlId++
	return "id-" + strconv.Itoa(htmlId)
}

func getWindow(view View) *Window {
	if view == nil {
		return nil
	}
	if w, ok := view.(*Window); ok {
		return w
	}
	return getWindow(view.parent())
}

func debugStr(value js.Value) string {
	sb := &strings.Builder{}
	sb.WriteString(value.Type().String())
	sb.WriteString(":")
	keys := js.Global().Get("Object").Call("keys", value)
	for i := 0; i < keys.Length(); i++ {
		sb.WriteString(keys.Index(i).String())
		sb.WriteString(",")
	}
	return sb.String()
}
