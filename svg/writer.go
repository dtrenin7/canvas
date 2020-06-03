package svg

import (
	"io"

	"github.com/dtrenin7/canvas"
)

// Writer writes the canvas as a SVG file
func Writer(w io.Writer, c *canvas.Canvas) error {
	svg := New(w, c.W, c.H)
	c.Render(svg)
	return svg.Close()
}
