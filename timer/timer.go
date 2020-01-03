// Package timer countdown duration
package timer

import (
	"math"
	"time"

	"github.com/pkg/errors"

	"github.com/gdamore/tcell"
	"github.com/hatappi/gomodoro/timer/screen"
)

// Timer interface
type Timer interface {
	Run(int) error
	Stop()
	IsQuit() bool

	ChangeFontColor(tcell.Color)
}

type timerImpl struct {
	title        string
	ticker       *time.Ticker
	screenClient screen.Client
	stopped      bool
	quit         bool

	fontColor      tcell.Color
	pauseFontColor tcell.Color
}

// NewTimer initilize Timer
func NewTimer(c screen.Client, title string) Timer {
	return &timerImpl{
		title:          title,
		ticker:         nil,
		screenClient:   c,
		fontColor:      tcell.ColorGreen,
		pauseFontColor: tcell.ColorDarkOrange,
	}
}

func (t *timerImpl) IsQuit() bool {
	return t.quit
}

func (t *timerImpl) ChangeFontColor(c tcell.Color) {
	t.fontColor = c
}

// Run timer
func (t *timerImpl) Run(duration int) error {
	drawFn := func(duration int, title string, opts ...screen.DrawOption) error {
		w, h := t.screenClient.ScreenSize()

		min := duration / 60
		sec := duration % 60

		x := float64(w) / 16
		y := float64(h) / 16

		printLine := 2.0
		cw := float64(w) * 14 / 16
		ch := float64(h) * 14 / 16
		ch -= printLine

		mag, err := getMagnification(cw, ch)
		if err != nil {
			return err
		}

		x = math.Round(x + ((cw - (screen.TimerWidth * mag)) / 2))
		y = math.Round(y + ((ch - (screen.TimerHeight * mag)) / 2))

		t.screenClient.Clear()
		t.screenClient.DrawSentence(int(x), int(y), int(screen.TimerWidth*mag), title)
		t.screenClient.DrawTimer(int(x), int(y)+2, int(mag), min, sec, opts...)

		return nil
	}

	err := drawFn(duration, t.title, screen.WithBackgroundColor(t.fontColor))
	if err != nil {
		return err
	}
	t.Start()
	for {
		opts := []screen.DrawOption{
			screen.WithBackgroundColor(t.fontColor),
		}
		select {
		case <-t.screenClient.GetQuitChan():
			t.quit = true
			return nil
		case <-t.screenClient.GetForceFinishChan():
			duration = 0
		case <-t.screenClient.GetPauseChan():
			if t.stopped {
				t.Start()
			} else {
				opts = []screen.DrawOption{
					screen.WithBackgroundColor(t.pauseFontColor),
				}
				t.Stop()
			}
		case <-t.ticker.C:
			duration--
		}

		err := drawFn(duration, t.title, opts...)
		if err != nil {
			return err
		}

		if duration == 0 {
			t.Stop()
			return nil
		}
	}
}

// Start timer
func (t *timerImpl) Start() {
	t.stopped = false
	t.ticker = time.NewTicker(1 * time.Second)
}

// Stop timer
func (t *timerImpl) Stop() {
	t.stopped = true
	t.ticker.Stop()
}

func getMagnification(w, h float64) (float64, error) {
	x := math.Round(w / screen.TimerWidth)
	y := math.Round(h / screen.TimerHeight)
	mag := math.Max(x, y)

	for {
		if mag < 1.0 {
			return 0.0, errors.New("screen is small")
		}

		if w >= screen.TimerWidth*mag && h >= screen.TimerHeight*mag {
			break
		}

		mag -= 1.0
	}

	return mag, nil
}
