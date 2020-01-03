package screen

import (
	"strings"

	"github.com/gdamore/tcell"
	runewidth "github.com/mattn/go-runewidth"
)

// DrawSentence is draw the sentence
func (c *clientImpl) DrawSentence(x, y, width int, str string, opts ...DrawOption) {
	str = adjustMessage(width, str)

	style := tcell.StyleDefault
	for _, opt := range opts {
		style = opt(style)
	}

	i := 0
	var deferred []rune
	dwidth := 0
	zwj := false
	for _, r := range str {
		if r == '\u200d' {
			if len(deferred) == 0 {
				deferred = append(deferred, ' ')
				dwidth = 1
			}
			deferred = append(deferred, r)
			zwj = true
			continue
		}
		if zwj {
			deferred = append(deferred, r)
			zwj = false
			continue
		}
		switch runewidth.RuneWidth(r) {
		case 0:
			if len(deferred) == 0 {
				deferred = append(deferred, ' ')
				dwidth = 1
			}
		case 1:
			if len(deferred) != 0 {
				c.screen.SetContent(x+i, y, deferred[0], deferred[1:], style)
				i += dwidth
			}
			deferred = nil
			dwidth = 1
		case 2:
			if len(deferred) != 0 {
				c.screen.SetContent(x+i, y, deferred[0], deferred[1:], style)
				i += dwidth
			}
			deferred = nil
			dwidth = 2
		}
		deferred = append(deferred, r)
	}
	if len(deferred) != 0 {
		c.screen.SetContent(x+i, y, deferred[0], deferred[1:], style)
	}

	c.screen.Show()
}

func adjustMessage(width int, str string) string {
	remain := (width - runewidth.StringWidth(str)) / 2
	if remain >= 0 {
		s := strings.Repeat(" ", remain)
		e := strings.Repeat(" ", width-runewidth.StringWidth(str)-remain)
		str = s + str + e
	} else {
		str = str[:width-3]
		str += "..."
	}

	return str
}
