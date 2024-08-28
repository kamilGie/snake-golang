package main

import (
	"fmt"
	"time"

	"github.com/kamilGie/snake-golang/snake"
	"github.com/kamilGie/snake-golang/snake/point"
	"github.com/nsf/termbox-go"
)

const ticksPerSecouns, width, higth = 3, 15, 10

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
  scoreText := fmt.Sprint(len(snakeBody) - 3)
	for index, char := range scoreText {
		termbox.SetCell(width+index+2, 1, char, termbox.ColorYellow, termbox.ColorDefault)
	}

	for _, value := range snakeBody {
		termbox.SetCell(value.X+1, value.Y+1, 'x', termbox.ColorGreen, termbox.ColorBlack)
	}
	termbox.SetCell(fruit.X+1, fruit.Y+1, 'f', termbox.ColorBlue, termbox.ColorBlack)
	termbox.Flush()
}

func handleInput(DirChan chan [4]int, endgame chan bool) {
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
			endgame <- true
		}
	}

}

func Update(snake *snake.Snake, DirChan chan [4]int, endGame chan bool) {
	for {
		select {
		case direction := <-DirChan:
			snake.TakeAction(direction)
		default:
			snake.TakeAction([4]int{0, 0, 0, 0})
		}
		snakeBody, fruit, gameOver := snake.GetState()
		if gameOver {
			endGame <- true
			return
		}
		DrawGame(snakeBody, fruit)
		time.Sleep(time.Second / ticksPerSecouns)
	}
}

func GameLoop() {
	snake := snake.New(width, higth)
	snakeDirectionChanel := make(chan [4]int)
	endGameChanel := make(chan bool)
	go handleInput(snakeDirectionChanel, endGameChanel)
	go Update(snake, snakeDirectionChanel, endGameChanel)
	<-endGameChanel
	return
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	GameLoop()
}
