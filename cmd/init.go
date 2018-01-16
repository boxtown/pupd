package cmd

import (
	"fmt"
	"log"

	"github.com/boxtown/pupd/model/pg"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
)

var initUser string
var initPass string

func init() {
	initCmd.Flags().StringVarP(&initUser, "user", "u", "postgres", "User to connect to the database with")
	initCmd.Flags().StringVarP(&initPass, "password", "p", "", "Password for connecting user")
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize the environment for PUPD",
	Run: func(cmd *cobra.Command, args []string) {
		url := fmt.Sprintf("postgres://%s:%s@localhost/pupd", initUser, initPass)
		source, err := pg.NewDataSource(url)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer source.Close()

		_, err = sqlx.LoadFile(source, "scripts/sql/schema.sql")
		if err != nil {
			log.Println(err.Error())
		}
	},
}
