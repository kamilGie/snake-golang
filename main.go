package main

import (
	"time"

	"github.com/kamilGie/snake-golang/snake"
	"github.com/kamilGie/snake-golang/snake/point"
	"github.com/nsf/termbox-go"
)

const width, higth = 15, 10

func DrawGame(snakeBody []point.Point, fruit point.Point) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for i := range higth + 1 {
		termbox.SetCell(0, i, '|', termbox.ColorWhite, termbox.ColorDefault)
		termbox.SetCell(width+1, i, '|', termbox.ColorWhite, termbox.ColorDefault)
	}
	for i := range width + 2 {
		termbox.SetCell(i, 0, '_', termbox.ColorWhite, termbox.ColorDefault)
		termbox.SetCell(i, higth+1, 'T', termbox.ColorWhite, termbox.ColorDefault)
	}
	for _, value := range snakeBody {
		termbox.SetCell(value.X+1, value.Y+1, 'x', termbox.ColorGreen, termbox.ColorBlack)
	}
	termbox.SetCell(fruit.X+1, fruit.Y+1, 'f', termbox.ColorBlue, termbox.ColorBlack)
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
		time.Sleep(time.Second / 3)
	}
}

func GameLoop() error {
	snake := snake.New(width, higth)
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
