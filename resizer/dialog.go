package main

import (
	"fmt"
	"image"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type Foo struct {
	Bar string
	Baz int
}

func sizeDialog(img image.Image) (uint, uint) {

	width := img.Bounds().Max.X - 1
	height := img.Bounds().Max.Y - 1

	// dialog window

	foo := &Foo{"b", 0}

	var valueEdit *walk.NumberEdit
	data := struct{ Value int }{100}

	MainWindow{
		Title:   "Quick Resize",
		MinSize: Size{300, 120},
		Layout:  VBox{},
		DataBinder: DataBinder{
			DataSource: foo,
			AutoSubmit: true,
			OnSubmitted: func() {
				fmt.Println(foo)
			},
		},
		Children: []Widget{
			// RadioButtonGroup is needed for data binding only.
			RadioButtonGroup{
				DataMember: "Bar",
				Buttons: []RadioButton{
					RadioButton{
						Name:  "aRB",
						Text:  "Pixel (width)",
						Value: "widthResize",
					},
					RadioButton{
						Name:  "bRB",
						Text:  "Pixel (height)",
						Value: "heightResize",
					},
					RadioButton{
						Name:  "cRB",
						Text:  "Percentage",
						Value: "percResize",
					},
				},
			},
			Composite{
				Layout:        Grid{Columns: 3},
				StretchFactor: 1,
				Children: []Widget{
					Label{Text: "Value: "},
					NumberEdit{
						AssignTo: &valueEdit,
						Value:    float64(data.Value),
						OnValueChanged: func() {
							data.Value = int(valueEdit.Value())
						},
					},
					PushButton{
						ColumnSpan: 3,
						Text:       "Resize!",
						Visible:    Bind("aRB.Checked"),
						OnClicked: func() {
							height = 0
							width = data.Value
						},
					},
					PushButton{
						ColumnSpan: 3,
						Text:       "Resize!",
						Visible:    Bind("bRB.Checked"),
						OnClicked: func() {
							height = data.Value
							width = 0
						},
					},
					PushButton{
						ColumnSpan: 3,
						Text:       "Resize!",
						Visible:    Bind("cRB.Checked"),
						OnClicked: func() {
							percHeight := data.Value / 100 * height
							height = percHeight
							width = 0
						},
					},
				},
			},
		},
	}.Run()

	// return values

	return uint(width), uint(height)

}