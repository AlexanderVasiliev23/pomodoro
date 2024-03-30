package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"time"
)

var finish time.Time
var progressBar *widget.ProgressBar
var seconds float64 = 5 * 60

func main() {

	myApp := app.New()
	w := myApp.NewWindow("Two Way")
	w.Resize(fyne.Size{
		Width:  300,
		Height: 200,
	})

	boundString := binding.NewString()
	if err := boundString.Set("0"); err != nil {
		panic(err)
	}
	progressBar = widget.NewProgressBar()

	finish = time.Now().Add(5 * time.Minute).Add(1 * time.Second)

	w.SetContent(container.NewVBox(
		container.NewVBox(progressBar),
		widget.NewEntryWithData(boundString),
		widget.NewButton("5", func() {
			finish = time.Now().Add(5 * time.Minute).Add(1 * time.Second)
		}),
		widget.NewButton("15", func() {
			finish = time.Now().Add(15 * time.Minute).Add(1 * time.Second)
		}),
		widget.NewButton("25", func() {
			finish = time.Now().Add(25 * time.Minute).Add(1 * time.Second)
		}),
	))

	go work(myApp, boundString)

	w.ShowAndRun()

}

func work(app fyne.App, boundString binding.String) {
	for {
		left := finish.Sub(time.Now())

		time.Sleep(100 * time.Millisecond)

		progressBar.SetValue(1 - left.Seconds()/seconds)

		msg := fmt.Sprintf("%02d:%02d", int(left.Minutes()), int(left.Seconds())%60)

		if err := boundString.Set(msg); err != nil {
			panic(err)
		}

		if left.Seconds() <= 0 {
			app.SendNotification(fyne.NewNotification("Pomodoro", "Finished"))
		}
	}
}
