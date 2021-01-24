package main

import (
	_ "fmt"
	"github.com/gdamore/tcell/v2"
	"log"
	"os"
	"time"
)

type Splash struct {
	screen *tcell.Screen
	Style  tcell.Style
	Note   string
}

type TimedSplash struct {
	Splash
	Timeout time.Duration
}

func NewTimedSplash(base tcell.Style, bg tcell.Color, note string, t time.Duration) *TimedSplash {
	return &TimedSplash{
		Timeout: t,
	}
}

func (e *Splash) Show(sc tcell.Screen, splash TimedSplash) {
	sc.Fill('x', splash.Style)
	sc.Show()
	time.Sleep(splash.Timeout)
}

func main() {
	screen, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Error creating new Screen: %v", err)
	}
	if err = screen.Init(); err != nil {
		log.Fatalf("Error initializing Screen: %v", err)
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

	splash := NewTimedSplash(screenBase, tcell.ColorRed, "Foopers", 200*time.Millisecond)

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
		}
		screen.Show()
	}
}
