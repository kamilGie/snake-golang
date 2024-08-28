package main

import (
	"github.com/kamilGie/snake-golang/snake"
	"github.com/kamilGie/snake-golang/snake/point"
	"github.com/nsf/termbox-go"
)

func DrawGame(snakeBody []point.Point, fruit point.Point) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	termbox.SetCell(fruit.X, fruit.Y, 'f', termbox.ColorBlue, termbox.ColorBlack)
	for _, value := range snakeBody {
		termbox.SetCell(value.X, value.Y, 'x', termbox.ColorGreen, termbox.ColorBlack)
	}
	termbox.Flush()
}

func GameLoop() error {
	snake := snake.New(10, 10)
	for {
		event := termbox.PollEvent()
		if event.Key == termbox.KeyArrowUp {
			snake.TakeAction([4]int{1, 0, 0, 0})
		} else if event.Key == termbox.KeyArrowLeft {
			snake.TakeAction([4]int{0, 1, 0, 0})
		} else if event.Key == termbox.KeyArrowRight {
			snake.TakeAction([4]int{0, 0, 0, 1})
		} else if event.Key == termbox.KeyArrowDown {
			snake.TakeAction([4]int{0, 0, 1, 0})
		} else if event.Key == termbox.KeyEsc {
			break
		}
		snakeBody, fruit, gameOver := snake.GetState()
		if gameOver {
			break
		}
		DrawGame(snakeBody, fruit)
	}
	return nil
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	err = GameLoop()
	if err != nil {
		panic(err)
	}

}
