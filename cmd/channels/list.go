package channels

import (
	"fmt"
	"os"

	"github.com/ohlmeier/youtube-cli/errors"
	"github.com/ohlmeier/youtube-cli/youtube"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list channels",
		Short: "List channels",
		Long:  ``,
		Run:   listChannels,
	}
)

func listChannels(ccmd *cobra.Command, args []string) {
	if len(args) > 0 {
		call := youtube.Service.Channels.List(args)
		call = call.ForUsername(args[0])
		response, err := call.Do()
		errors.Handle(err, "Error calling youtube api to display channels")
		// TODO: add more channel infos
		for _, item := range response.Items {
			fmt.Printf("ID:%s\tTitle:%s\tViews%d\n", item.Id, item.Snippet.Title, item.Statistics.ViewCount)
		}
	} else {
		fmt.Fprintln(os.Stderr, "No username is specified. Please specify a youtube username to display the channels.")
		return
	}
}
