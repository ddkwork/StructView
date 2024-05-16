package main

import (
	"StructView"

	"github.com/ddkwork/app"
	"github.com/richardwilkes/unison"
)

func main() {
	app.Run("StructView", func(w *unison.Window) {
		StructView.New(w).Layout(w.Content())
	})
}
