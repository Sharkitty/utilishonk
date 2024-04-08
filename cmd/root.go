package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"utilishonk/cmd/distrobox"
)

var rootCmd = &cobra.Command{
	Use:   "shonk",
	Short: "Utilishonk makes shonks hapi with CLI utilities",
	Long:  "Utilishonk is a set of CLI utilities written in Go, that makes shonks life easier",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("This is the main command")
	},
}

func init() {
	rootCmd.AddCommand(distrobox.DistroboxCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
