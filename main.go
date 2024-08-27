package main

import (
	"fmt"
	"time"
  "github.com/kamilGie/snake-golang/snake"
	"github.com/nsf/termbox-go"
)

func GameLoop() error {
	ticker := time.NewTicker(time.Second / 10)
  defer ticker.Stop()
  i := 0
  s := snake.Snake{}
	for { 
    fmt.Println(i)
    i++
		event := termbox.PollEvent()
		if event.Key == termbox.KeyArrowUp {
		} else if event.Key == termbox.KeyArrowLeft {
		} else if event.Key == termbox.KeyArrowRight {
		} else if event.Key == termbox.KeyArrowDown {
		} else if event.Key == termbox.KeyEsc {
			break
		}
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
