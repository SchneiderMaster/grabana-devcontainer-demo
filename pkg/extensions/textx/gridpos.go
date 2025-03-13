package textx

import "github.com/K-Phoen/grabana/text"

func GridPos(h, w, x, y int) text.Option {
	return func(text *text.Text) error {
		text.Builder.GridPos.H = &h
		text.Builder.GridPos.W = &w
		text.Builder.GridPos.X = &x
		text.Builder.GridPos.Y = &y
		return nil
	}
}
