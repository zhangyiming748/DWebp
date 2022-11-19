package DWebp

import "testing"

func TestUnit(t *testing.T) {
	src := "/Users/zen/Github/CWebp/mac"
	dst := "/Users/zen/Github/DWebp/done"
	patttern := "png"
	DWebp(src, dst, patttern)
}

func TestDWebM2Gif(t *testing.T) {
	src := "/Users/zen/Pictures/阿尼亚表情/gif"
	dst := "/Users/zen/Pictures/阿尼亚表情"
	patttern := "gif"
	DWebM2Gif(src, dst, patttern)
}
