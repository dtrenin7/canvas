// +build gofuzz
package fuzz

import "github.com/dtrenin7/canvas/font"

func Fuzz(data []byte) int {
	_, _ = font.ParseWOFF2(data)
	return 1
}
