package etl

import "context"

type Input struct {
}

type Output struct {
}

type Extractor interface {
	Extract(context.Context, chan<- Input) error
}

type Transformer interface {
	Transform(context.Context, <-chan Input, chan<- Output) error
}

type Loader interface {
	Load(context.Context, <-chan Output) error
}
