package extract

import (
	"../etl"
	"context"
)

var _ etl.Extractor = &SongExtractor{}

type SongExtractor struct {
}

func (e *SongExtractor) Extract(ctx context.Context, input chan<- etl.Input) error {

	return nil
}
