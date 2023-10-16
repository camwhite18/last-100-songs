package load

import (
	"../etl"
	"context"
)

var _ etl.Loader = &SpotifyLoader{}

type SpotifyLoader struct {
}

func (l *SpotifyLoader) Load(ctx context.Context, input <-chan etl.Output) error {

	return nil
}
