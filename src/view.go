package tbp

import termbox "github.com/nsf/termbox-go"

// Colors
const boardColor = termbox.ColorDefault
const backgroundColor = termbox.ColorDefault
const instructionsColor = termbox.ColorWhite

// Layout
const defaultMarginWidth = 0
const defaultMarginHeight = 1
const titleStartX = defaultMarginWidth
const titleStartY = defaultMarginHeight
const titleHeight = 1
const titleEndY = titleStartY + titleHeight
const boardStartX = defaultMarginWidth
const boardStartY = titleEndY + defaultMarginHeight
const boardWidth = 240
const boardHeight = 70
const cellWidth = 2
const boardEndX = boardStartX + boardWidth*cellWidth
const boardEndY = boardStartY + boardHeight
const instructionsStartX = boardEndX + defaultMarginWidth
const instructionsStartY = boardStartY

// Text in the UI
const title = "THREE BODY PROBLEM"

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

	termbox.SetCell(xmiddle, ymiddle, '+', termbox.ColorWhite, termbox.ColorDefault)
	for idx, body := range bodies {
		xpos := boardStartX + cellWidth*(int(body.Position.x/unzoom)) + xmiddle
		ypos := boardStartY + (int(body.Position.y / unzoom)) + ymiddle

		// println("===========================")
		// println(body.Position.x)
		// println(body.Position.y)

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
