package transform

import (
	"../etl"
	"context"
)

var _ etl.Transformer = &Transformer{}

type Transformer struct {
}

func (t *Transformer) Transform(ctx context.Context, input <-chan etl.Input, output chan<- etl.Output) error {

	return nil
}
