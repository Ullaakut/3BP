package tbp

import termbox "github.com/nsf/termbox-go"

// Colors
const boardColor = termbox.ColorDefault
const backgroundColor = termbox.ColorDefault
const instructionsColor = termbox.ColorWhite

// Layout
const boardWidth = 500
const boardHeight = 70
const cellWidth = 2

// Text in the UI
const title = "THREE BODY PROBLEM"

const unzoom = 1

var cellColors = []termbox.Attribute{
	termbox.ColorRed,
	termbox.ColorGreen,
	termbox.ColorBlue,
}

// Render takes care of rendering everything.
func Render(bodies []Body) {
	termbox.Clear(backgroundColor, backgroundColor)
	xmiddle := boardWidth / (2 * cellWidth)
	ymiddle := boardHeight / 2

	// averageposx := (bodies[0].Position.x + bodies[1].Position.x + bodies[2].Position.x) / 3
	// averageposy := (bodies[0].Position.y + bodies[1].Position.y + bodies[2].Position.y) / 3

	termbox.SetCell(xmiddle, ymiddle, '+', termbox.ColorWhite, termbox.ColorDefault)
	for idx, body := range bodies {
		xpos := int(float64(cellWidth)*(body.Position.x/float64(unzoom)) + float64(xmiddle))
		ypos := int((body.Position.y / float64(unzoom)) + float64(ymiddle))

		// Follow bodies through space
		// xpos = xpos - averageposx
		// ypos = ypos - averageposy

		if xpos > 0 &&
			xpos < boardWidth*cellWidth &&
			ypos > 0 &&
			ypos < boardHeight {

			for i := 0; i < cellWidth; i++ {
				color := cellColors[idx]
				if xpos+i != xmiddle || ypos+i != ymiddle {
					termbox.SetCell(xpos+i, ypos, ' ', termbox.ColorWhite, color)
				} else {
					termbox.SetCell(xpos+i, ypos, '+', termbox.ColorWhite, color)
				}
			}
		}
	}

	termbox.Flush()
}
