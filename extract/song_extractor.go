package extract

import (
	"../etl"
	"context"
)

var _ etl.Extractor = &SongExtractor{}

type SongExtractor struct {
	stationURL string
}

func NewSongExtractor(stationURL string) *SongExtractor {
	return &SongExtractor{
		stationURL: stationURL,
	}
}

func (e *SongExtractor) Extract(ctx context.Context, input chan<- etl.Input) error {

	return nil
}
