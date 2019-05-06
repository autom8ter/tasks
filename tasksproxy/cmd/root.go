package cmd

import (
	"fmt"
	"github.com/autom8ter/tasks/sdk/go/tasks"
	"google.golang.org/grpc"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const apiAddr = "104.198.16.50:8080"

var client tasks.TaskServiceClient

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "tasksproxy",
	Long: `
Tasksproxy
`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	conn, err := grpc.Dial(apiAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err.Error())
	}
	client = tasks.NewTaskServiceClient(conn)
}
