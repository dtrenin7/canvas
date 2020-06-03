package pdf

import (
	"io"

	"github.com/dtrenin7/canvas"
)

// Writer writes the canvas as a PDF file.
func Writer(w io.Writer, c *canvas.Canvas) error {
	pdf := New(w, c.W, c.H)
	c.Render(pdf)
	return pdf.Close()
}
