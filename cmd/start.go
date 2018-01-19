package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/boxtown/pupd/api"
	"github.com/boxtown/pupd/model/pg"
	"github.com/spf13/cobra"
)

var apiDBUser string
var apiDBPass string

func init() {
	startCmd.Flags().StringVarP(&apiDBUser, "user", "u", "postgres", "User to connect to the database with")
	startCmd.Flags().StringVarP(&apiDBPass, "password", "p", "", "Password for connecting user")
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the PUPD API server",
	Run: func(cmd *cobra.Command, args []string) {
		connStr := fmt.Sprintf("postgres://%s:%s@localhost/pupd", apiDBUser, apiDBPass)
		source, err := pg.NewDataSource(connStr)
		if err != nil {
			log.Fatal(err)
		}
		defer source.Close()

		routes := api.Router(source)
		s := &http.Server{Addr: ":3000", Handler: routes}

		stop := make(chan os.Signal, 1)
		errc := make(chan error, 1)
		signal.Notify(stop, os.Interrupt)
		go func(errc chan<- error) {
			log.Println("Starting server on :3000...")
			if err := s.ListenAndServe(); err != nil {
				errc <- err
			}
		}(errc)

		for {
			select {
			case <-stop:
				log.Println("Shutting down server...")
				s.Shutdown(context.Background())
				return
			case err := <-errc:
				log.Println(err)
				return
			}
		}
	},
}
