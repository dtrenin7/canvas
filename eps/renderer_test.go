package eps

import (
	"bytes"
	"testing"

	"github.com/dtrenin7/canvas"
)

func TestEPS(t *testing.T) {
	w := &bytes.Buffer{}
	eps := New(w, 100, 80)
	eps.setColor(canvas.Red)
	//test.String(t, string(w.Bytes()), "")
}
