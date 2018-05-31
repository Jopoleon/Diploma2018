package main

import (
	"strconv"

	"fmt"

	"github.com/andlabs/ui"
)

func main() {
	err := ui.Main(func() {
		input := ui.NewEntry()
		p0 := ui.NewEntry()
		T0 := ui.NewEntry()
		u := ui.NewEntry()
		button := ui.NewButton("Greet")
		greeting := ui.NewLabel("")
		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Enter start params:"), false)
		box.Append(input, false)
		box.Append(p0, false)
		box.Append(T0, false)
		box.Append(u, false)
		box.Append(button, false)
		box.Append(greeting, false)
		window := ui.NewWindow("Hello", 300, 200, false)
		window.SetMargined(true)
		window.SetChild(box)
		Abutton := ui.NewButton("***Advanced options***")
		Abutton.OnClicked(func(*ui.Button) {
			ab := ui.NewVerticalBox()
			//Awindow := ui.NewWindow("Advanced options", 200, 100, false)
			//window.
			ab.Append(ui.NewLabel("a1"), false)
			a1 := ui.NewEntry()
			ab.Append(a1, false)
			ab.Append(ui.NewLabel("a2"), false)
			a2 := ui.NewEntry()
			ab.Append(a2, false)
			ab.Append(ui.NewLabel("a3"), false)
			a3 := ui.NewEntry()
			ab.Append(a3, false)
			ab.Append(ui.NewLabel("a4"), false)
			a4 := ui.NewEntry()
			ab.Append(a4, false)
			b := ui.NewButton("Submit advanced params")
			ab.Append(b, false)
			b.OnClicked(func(*ui.Button) {
				//window.Show()
				fmt.Println(a1.Text(), a2.Text(), a3.Text(), a4.Text())
				window.SetChild(box)
				//window.Show()
			})

			window.SetChild(ab)
			//greeting.SetText("Hello, " + input.Text() + "!")
		})
		box.Append(Abutton, false)
		button.OnClicked(func(*ui.Button) {
			_, _, _, _, err := ParseParams(input.Text(), p0.Text(), T0.Text(), u.Text())
			if err != nil {
				greeting.SetText("Bad param: " + err.Error())
			}
			//greeting.SetText("Hello, " + input.Text() + "!")
		})
		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}

func ParseParams(p0s, T0s, us, somes string) (p0, T0, u, some float64, err error) {
	//var err error
	p0, err = strconv.ParseFloat(p0s, 64)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("%s", "Param p0 is in invalid format")
	}
	T0, err = strconv.ParseFloat(T0s, 64)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("%s", "Param T0 is in invalid format")
	}
	u, err = strconv.ParseFloat(us, 64)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("%s", "Param T0 is in invalid format")
	}
	some, err = strconv.ParseFloat(somes, 64)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("%s", "Param some is in invalid format")
	}
	fmt.Println(p0, T0, u, some)
	return p0, T0, u, some, nil
}
