package cmd

import (
	"fmt"

	"github.com/HashedDan/charon/pkg/server"
	"github.com/spf13/cobra"
)

var port string

func init() {
	startCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to serve charon on.")
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the charon server.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("charon starting...")
		srv := server.NewHttpServer(port)
		srv.Start()
	},
}
