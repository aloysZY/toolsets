package myTime

import (
	"github.com/aloysZY/MyOperatorProjects/toolsets/internak/time/now"
	"github.com/spf13/cobra"
)

func init() {
	TimeCmd.AddCommand(now.NowCmd)
}

var TimeCmd = &cobra.Command{
	Use:   "time",
	Short: "Get the status of etcd cluster member",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
