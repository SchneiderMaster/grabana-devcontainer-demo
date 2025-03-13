package timeseriesx

import "github.com/K-Phoen/grabana/timeseries"

func GridPos(h, w, x, y int) timeseries.Option {
	return func(text *timeseries.TimeSeries) error {
		text.Builder.GridPos.H = &h
		text.Builder.GridPos.W = &w
		text.Builder.GridPos.X = &x
		text.Builder.GridPos.Y = &y
		return nil
	}
}
