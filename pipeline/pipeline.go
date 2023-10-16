package pipeline

import (
	"../etl"
	"context"
)

type Pipeline struct {
	Extractor   etl.Extractor
	Transformer etl.Transformer
	Loader      etl.Loader
}

func (p *Pipeline) Run(ctx context.Context) error {
	numOfSteps := 3
	input := make(chan etl.Input, 1)
	output := make(chan etl.Output, 1)
	errChan := make(chan error, numOfSteps)

	go func() {
		defer close(input)
		errChan <- p.Extractor.Extract(ctx, input)
	}()

	go func() {
		defer close(output)
		errChan <- p.Transformer.Transform(ctx, input, output)
	}()

	go func() {
		errChan <- p.Loader.Load(ctx, output)
	}()

	for i := 0; i < numOfSteps; i++ {
		if err := <-errChan; err != nil {
			return err
		}
	}

	return nil
}
