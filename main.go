package main

import (
	"time"

	"github.com/nsf/termbox-go"
)

func GameLoop() error {
	ticker := time.NewTicker(time.Second / 10)
	for range ticker.C {
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

	GameLoop()
	if err != nil {
		panic(err)
	}

}
