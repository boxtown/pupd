package cmd

import (
	"fmt"
	"log"

	"github.com/boxtown/pupd/model"
	"github.com/boxtown/pupd/model/pg"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/cobra"
)

var seedUser string
var seedPass string

func init() {
	seedCmd.Flags().StringVarP(&seedUser, "user", "u", "postgres", "User to connect to the database with")
	seedCmd.Flags().StringVarP(&seedPass, "password", "p", "", "Password for connecting user")
	rootCmd.AddCommand(seedCmd)
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed the database with test data",
	Run: func(cmd *cobra.Command, args []string) {
		movements := []model.Movement{
			model.Movement{Name: "Test"},
		}

		url := fmt.Sprintf("postgres://%s:%s@localhost/pupd", seedUser, seedPass)
		db, err := sqlx.Connect("postgres", url)
		if err != nil {
			log.Fatal(err.Error())
		}

		movementStore := pg.NewMovementStore(db)
		for _, movement := range movements {
			_, err := movementStore.Create(&movement)
			if err != nil {
				db.Close()
				log.Fatal(err)
			}
		}

		db.Close()
	},
}
