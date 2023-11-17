package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yt",
	Short: "A YouTube CLI.",
	Long:  `This CLI can be used to access the YouTube API.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is a YouTube CLI.")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
