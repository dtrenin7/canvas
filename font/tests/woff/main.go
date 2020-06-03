// +build gofuzz
package fuzz

import "github.com/dtrenin7/canvas/font"

func Fuzz(data []byte) int {
	_, _ = font.ParseWOFF(data)
	return 1
}
