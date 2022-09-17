package main

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

var CIRCLE_SIZE = fyne.NewSize(5, 5)

func initBaseCircles() [6]*canvas.Circle {
	var circles [6]*canvas.Circle
	circles[0] = canvas.NewCircle(color.Black)
	circles[0].Move(fyne.NewPos(225, 100))
	circles[0].Resize(CIRCLE_SIZE)

	circles[1] = canvas.NewCircle(color.Black)
	circles[1].Move(fyne.NewPos(525, 100))
	circles[1].Resize(CIRCLE_SIZE)

	circles[2] = canvas.NewCircle(color.Black)
	circles[2].Move(fyne.NewPos(75, 350))
	circles[2].Resize(CIRCLE_SIZE)

	circles[3] = canvas.NewCircle(color.Black)
	circles[3].Move(fyne.NewPos(675, 350))
	circles[3].Resize(CIRCLE_SIZE)

	circles[4] = canvas.NewCircle(color.Black)
	circles[4].Move(fyne.NewPos(225, 600))
	circles[4].Resize(CIRCLE_SIZE)

	circles[5] = canvas.NewCircle(color.Black)
	circles[5].Move(fyne.NewPos(525, 600))
	circles[5].Resize(CIRCLE_SIZE)

	return circles

}

func main() {
	a := app.New()
	w := a.NewWindow("Drawn")
	circles := initBaseCircles()

	base := container.NewWithoutLayout()
	for _, circle := range circles {
		base.AddObject(circle)
	}

	w.SetContent(base)

	w.Resize(fyne.NewSize(800, 800))

	go drawLoop(circles, base)

	w.ShowAndRun()
}

func drawLoop(circles [6]*canvas.Circle, base *fyne.Container) {
	x, y := circles[0].Position().X, circles[0].Position().Y
	for {
		poss := rand.Intn(6)
		targetCircle := circles[poss]
		targetX, targetY := targetCircle.Position().X, targetCircle.Position().Y
		if x != targetX {
			sub := math.Abs(float64(x - targetX))
			twoThirds := sub / 3 * 2
			if x > targetX {
				x = float32(math.Abs(float64(x) - twoThirds))
			} else {
				x = float32(float64(x) + twoThirds)
			}
		}
		if y != targetY {
			sub := math.Abs(float64(y - targetY))
			twoThirds := sub / 3 * 2
			if y > targetY {
				y = float32(math.Abs(float64(y) - twoThirds))
			} else {
				y = float32(float64(y) + twoThirds)
			}
		}
		circle := canvas.NewCircle(color.Black)
		circle.Move(fyne.NewPos(x, y))
		circle.Resize(CIRCLE_SIZE)
		base.AddObject(circle)
		base.Refresh()
		time.Sleep(time.Microsecond * 500)
	}
}
