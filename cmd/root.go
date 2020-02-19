package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "falcon",
	Short: "Falcon is fully auto-configure package manager for C/C++ open source projects",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("Falcon executed xD")
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)
}
