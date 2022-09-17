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

func createCircle(position fyne.Position) *canvas.Circle {
	circle := canvas.NewCircle(color.Black)
	circle.Move(position)
	circle.Resize(fyne.NewSize(4, 4))
	return circle
}

func initBaseCircles() [6]*canvas.Circle {
	var circles [6]*canvas.Circle
	circles[0] = createCircle(fyne.NewPos(225, 100))
	circles[1] = createCircle(fyne.NewPos(525, 100))
	circles[2] = createCircle(fyne.NewPos(75, 350))
	circles[3] = createCircle(fyne.NewPos(675, 350))
	circles[4] = createCircle(fyne.NewPos(225, 600))
	circles[5] = createCircle(fyne.NewPos(525, 600))

	return circles
}

func calculatePosition(initial float32, target float32) float32 {
	if initial != target {
		sub := math.Abs(float64(initial - target))
		twoThirds := sub / 3 * 2
		if initial > target {
			return float32(float64(initial) - twoThirds)
		} else {
			return float32(float64(initial) + twoThirds)
		}
	}
	return initial
}

func drawLoop(circles [6]*canvas.Circle, base *fyne.Container) {
	x, y := circles[0].Position().X, circles[0].Position().Y
	for {
		poss := rand.Intn(6)
		targetCircle := circles[poss]
		targetX, targetY := targetCircle.Position().X, targetCircle.Position().Y
		x = calculatePosition(x, targetX)
		y = calculatePosition(y, targetY)
		circle := createCircle(fyne.NewPos(x, y))
		base.Add(circle)
		base.Refresh()
		time.Sleep(time.Microsecond * 500)
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Drawn")
	circles := initBaseCircles()

	base := container.NewWithoutLayout()
	for _, circle := range circles {
		base.Add(circle)
	}

	w.SetContent(base)

	w.Resize(fyne.NewSize(800, 800))

	go drawLoop(circles, base)

	w.ShowAndRun()
}
