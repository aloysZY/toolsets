package now

import (
	"github.com/spf13/cobra"
)

var NowCmd = &cobra.Command{
	Use:   "now",
	Short: "Get the status of etcd cluster member",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
