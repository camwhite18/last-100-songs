package load

import (
	"../etl"
	"context"
)

var _ etl.Loader = &SpotifyLoader{}

type SpotifyLoader struct {
	ClientID     string
	ClientSecret string
	AccessToken  string
}

func NewSpotifyLoader(clientID string, clientSecret string) *SpotifyLoader {
	// Make a request to Spotify to get an access token if the current one is expired

	return &SpotifyLoader{
		AccessToken: accessToken,
	}
}

func (l *SpotifyLoader) Load(ctx context.Context, input <-chan etl.Output) error {

	return nil
}
