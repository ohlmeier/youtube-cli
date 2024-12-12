package youtube

import (
	"log"
	"os"

	"github.com/ohlmeier/youtube-cli/errors"
	"golang.org/x/net/context"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

var Service *youtube.Service

func init() {
	ctx := context.Background()
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatal("API Key missing")
	}
	youtubeService, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	errors.Handle(err, "Error creating YouTube service")
	Service = youtubeService
}
