package main

import (
	"StructView"

	"github.com/ddkwork/app/widget"

	"github.com/ddkwork/unison"
	"github.com/ddkwork/unison/app"
)

func main() {
	app.Run("Keygen for data recovery", func(w *unison.Window) {
		v := structView.New().Layout()
		panel := widget.NewPanel()
		panel.AddChild(v.View)
		panel.AddChild(widget.NewVSpacer())
		panel.AddChild(v.RowPanel)
		scrollPanelFill := widget.NewScrollPanelFill(panel)
		w.Content().AddChild(scrollPanelFill)
	})
}
