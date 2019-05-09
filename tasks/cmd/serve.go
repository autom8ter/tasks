package cmd

import (
	"fmt"
	"github.com/autom8ter/tasks/pkg/config"
	"github.com/autom8ter/tasks/pkg/functions"
	"github.com/autom8ter/tasks/pkg/service"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var dbCfg = &config.DBConfig{}
var port int

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "start the tasks server",
	Run: func(cmd *cobra.Command, args []string) {
		svc, err := service.NewService(
			functions.WithDBConfig(dbCfg),
		)
		if err != nil {
			log.Fatalln(err.Error())
		}
		if err := svc.Serve(fmt.Sprintf(":%v", port)); err != nil {
			log.Fatalln(err.Error())
		}
	},
}

func init() {
	_ = godotenv.Load()
	serveCmd.Flags().IntVar(&port, "port", 8080, "port to serve on")
	serveCmd.Flags().StringVarP(&dbCfg.Database, "db", "d", "tasks", "postgres database name")
	serveCmd.Flags().StringVarP(&dbCfg.User, "user", "u", "postgres", "postgres user name")
	serveCmd.Flags().StringVarP(&dbCfg.Password, "pass", "p", os.Getenv("POSTGRES_PASSWORD"), "postgres password")
	serveCmd.Flags().StringVarP(&dbCfg.Addr, "addr", "a", os.Getenv("POSTGRES_ADDRESS"), "postgres remote address")

	rootCmd.AddCommand(serveCmd)

}
