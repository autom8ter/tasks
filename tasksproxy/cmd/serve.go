package cmd

import (
	"fmt"
	"github.com/autom8ter/tasks/pkg/functions"
	"github.com/autom8ter/tasks/pkg/proxy"
	"github.com/spf13/cobra"
	"log"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start the tasks proxy server",
	Run: func(cmd *cobra.Command, args []string) {
		proxyy = proxy.NewProxy(
			functions.WithProfiling(),
			functions.WithMetrics(),
			functions.WithDocs(),
			functions.WithTaskCreate(client),
			functions.WithTaskDelete(client),
			functions.WithTaskList(client),
			functions.WithTaskRead(client),
			functions.WithTaskUpdate(),
			functions.WithTaskList(client),
		)
		fmt.Println("starting proxy server on http://localhost:8080")
		err := proxyy.ListenAndServe(":8080")
		if err != nil {
			log.Println(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
