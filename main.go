package main

import (
	//"fmt"
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
	x := 0
	for {
		select {
		case <-abort:
			//splash.Show(screen, *splash)
			screen.Fill('x', splash.Style)
			screen.Show()
			time.Sleep(splash.Timeout)
			x++
			return
		case <-tick:
			var s tcell.Style
			s = s.Background(tcell.ColorGreen)
			//s = tcell.Style(tcell.ColorGreen)
			//s = s.Background(tcell.ColorGreen)
			screen.Fill('x', s)
			//screenBase.Background(splash)
			p := screenBase.Background(tcell.ColorPurple)
			screen.SetContent(x, 1, []rune{'a', 'c', 'd'}, p)
			screen.Show()
		}
		screen.Show()
	}
}
