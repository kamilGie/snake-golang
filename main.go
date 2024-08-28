package main

import (
	"time"

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

func handleInput(DirChan chan [4]int) {
	for {
		event := termbox.PollEvent()
		if event.Key == termbox.KeyArrowUp {
			DirChan <- [4]int{1, 0, 0, 0}
		} else if event.Key == termbox.KeyArrowLeft {
			DirChan <- [4]int{0, 1, 0, 0}
		} else if event.Key == termbox.KeyArrowRight {
			DirChan <- [4]int{0, 0, 0, 1}
		} else if event.Key == termbox.KeyArrowDown {
			DirChan <- [4]int{0, 0, 1, 0}
		} else if event.Key == termbox.KeyEsc {
			return
		}
	}

}

func Update(snake *snake.Snake, DirChan chan [4]int) {
	for {
		select {
		case direction := <-DirChan:
			snake.TakeAction(direction)
		default:
			snake.TakeAction([4]int{0, 0, 0, 0})
		}
		snakeBody, fruit, gameOver := snake.GetState()
		if gameOver {
			return
		}
		DrawGame(snakeBody, fruit)
		time.Sleep(time.Second)
	}
}

func GameLoop() error {
	snake := snake.New(10, 10)
	snakeDirectionChanel := make(chan [4]int)
	go func() {
		handleInput(snakeDirectionChanel)
	}()
	Update(snake, snakeDirectionChanel)
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
