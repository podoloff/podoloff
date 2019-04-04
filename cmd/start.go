package cmd

import (
	"fmt"

	"github.com/hasheddan/charon/pkg/server"
	"github.com/spf13/cobra"
)

var port, db string

func init() {
	startCmd.Flags().StringVarP(&port, "port", "p", "8080", "Port to serve charon on.")
	startCmd.Flags().StringVarP(&db, "database", "d", "mongodb://localhost:27017", "Connection string for Mongo.")
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the charon server.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("charon starting...")
		srv := server.NewHTTPServer(port, db)
		srv.Start()
	},
}
