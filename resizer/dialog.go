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

type MyMainWindow struct {
	*walk.MainWindow
}

func sizeDialog(img image.Image, filetype string) (uint, uint, bool) {

	width := img.Bounds().Max.X - 1
	height := img.Bounds().Max.Y - 1

	var savepng bool

	// dialog window

	var appIcon, _ = walk.NewIconFromResourceId(3)

	mw := new(MyMainWindow)

	foo := &Foo{"widthResize", 0}

	var valueEdit *walk.NumberEdit
	data := struct{ Value int }{width}

	MainWindow{
		AssignTo: &mw.MainWindow,
		Title:    "Quick Resize",
		Icon:     appIcon,
		MinSize:  Size{300, 120},
		Layout:   VBox{},
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
					}, /*
						RadioButton{
							Name:  "cRB",
							Text:  "Percentage",
							Value: "percResize",
						}, */
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
					CheckBox{
						Name:    "saveaspng",
						Text:    "Save as .png",
						Checked: true,
						Visible: func() bool {
							if filetype == "jpg" {
								return true
							} else {
								return false
							}
						},
					},
					PushButton{
						ColumnSpan: 3,
						Text:       "Resize!",
						Visible:    Bind("aRB.Checked && saveaspng.Checked"),
						OnClicked: func() {
							height = 0
							width = data.Value
							var savepng = true
							if savepng {
							} else {
							}
							mw.Close()
						},
					},
					PushButton{
						ColumnSpan: 3,
						Text:       "Resize!",
						Visible:    Bind("bRB.Checked && saveaspng.Checked"),
						OnClicked: func() {
							height = data.Value
							width = 0
							var savepng = true
							if savepng {
							} else {
							}
							mw.Close()
						},
					},
					PushButton{
						ColumnSpan: 3,
						Text:       "Resize!",
						Visible:    Bind("aRB.Checked && !saveaspng.Checked"),
						OnClicked: func() {
							height = 0
							width = data.Value
							var savepng = false
							if savepng {
							} else {
							}
							mw.Close()
						},
					},
					PushButton{
						ColumnSpan: 3,
						Text:       "Resize!",
						Visible:    Bind("bRB.Checked && !saveaspng.Checked"),
						OnClicked: func() {
							height = data.Value
							width = 0
							var savepng = false
							if savepng {
							} else {
							}
							mw.Close()
						},
					}, /*
						PushButton{
							ColumnSpan: 3,
							Text:       "Resize!",
							Visible:    Bind("cRB.Checked"),
							OnClicked: func() {
								percHeight := float64(data.Value) / 100
								height = int(percHeight) * height
								width = 0
							},
						}, */
				},
			},
		},
	}.Run()

	// return values

	return uint(width), uint(height), savepng

}
