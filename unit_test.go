package DWebp

import "testing"

func TestUnit(t *testing.T) {
	src := "/Users/zen/Github/CWebp/mac"
	dst := "/Users/zen/Github/DWebp/done"
	patttern := "png"
	DWebp(src, dst, patttern)
}
