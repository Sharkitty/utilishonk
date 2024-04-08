package distrobox

import (
	"fmt"

	"github.com/spf13/cobra"
)

var DistroboxCmd = &cobra.Command{
	Use:   "distrobox",
	Short: "Distrobox utilities",
	Long:  "Use distrobox related utilities",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Omg it's the distrobox util!")
	},
}
