package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
	"os"
	"time"
)

type ExitSplash struct {
	Style    tcell.Style
	Note     string
	Lifetime time.Duration
}

func NewSplash(base tcell.Style, bg tcell.Color, note string) *ExitSplash {
	return &ExitSplash{
		Style:    base.Background(bg),
		Note:     note,
		Lifetime: 500 * time.Millisecond,
	}
}

func (e *ExitSplash) Show(sc tcell.Screen, splash ExitSplash) {
	sc.Fill('x', splash.Style)
	sc.Show()
	time.Sleep(splash.Lifetime)
}

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		fmt.Errorf("Error creating new Screen: %v", err)
	}
	if err = screen.Init(); err != nil {
		fmt.Errorf("Error initializing Screen: %v", err)
	}
	// TODO callback to finish up
	defer screen.Fini()

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	//var screenBase tcell.Style
	screenBase := tcell.StyleDefault

	splash := NewSplash(screenBase, tcell.ColorRed, "Foopers")

	tick := time.Tick(100 * time.Millisecond)
	for {
		select {
		case <-abort:
			splash.Show(screen, *splash)
			return
		case <-tick:
			// screen.Fill('x', tcell.Style{tcell.ColorGreen})
			screen.SetContent(10, 10, 'a', []rune{}, tcell.StyleDefault)
			screen.Show()
		default:
			//
		}
		screen.Show()
	}
}
