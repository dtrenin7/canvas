// +build gofuzz
package fuzz

import "github.com/dtrenin7/canvas"

func Fuzz(data []byte) int {
	_, _ = canvas.ParseLaTeX(string(data))
	return 1
}
