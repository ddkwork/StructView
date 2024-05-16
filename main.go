package main

import (
	"StructView"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("StructView", func(w *unison.Window) {
		structView.New(w).Layout(w.Content())
	})
}
