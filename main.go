package main

import (
	"github.com/nsf/termbox-go"
	"github.com/ullaakut/3BD/src"
)

func main() {
	bodies := []tbp.Body{
		tbp.Body1,
		tbp.Body2,
		tbp.Body3,
	}

	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	for {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch {
				case ev.Ch == 'q' || ev.Key == termbox.KeyEsc || ev.Key == termbox.KeyCtrlC || ev.Key == termbox.KeyCtrlD:
					return
				}
			}
		default:
			tbp.Render(bodies)
			bodies = tbp.ProcessBodies(bodies, 10, 10)
		}
	}
}
