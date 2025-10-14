package cmd

import "github.com/spf13/cobra"

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start the HTTP server",
	Run: func(cmd *cobra.Command, args []string) {
		// app.RunServer()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
