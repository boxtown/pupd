package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

// Seed is a structure used to represent a database seed data store
type Seed struct {
	Movements []*model.Movement `json:"movements"`
	Workouts  []*model.Workout  `json:"workouts"`
}

var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "Seed the database with test data",
	Run: func(cmd *cobra.Command, args []string) {
		var seed Seed
		raw, err := ioutil.ReadFile("./resources/seed.json")
		if err != nil {
			log.Fatal(err)
		}
		err = json.Unmarshal(raw, &seed)
		if err != nil {
			log.Fatal(err)
		}

		url := fmt.Sprintf("postgres://%s:%s@localhost/pupd", seedUser, seedPass)
		source, err := pg.NewDataSource(url)
		if err != nil {
			log.Fatal(err)
		}
		defer source.Close()

		err = source.Transaction(func(tx *sqlx.Tx) error {
			movementStore := pg.NewMovementStore(tx)
			for _, movement := range seed.Movements {
				id, err := movementStore.Create(movement)
				if err != nil {
					return err
				}
				movement.ID = id
			}

			workoutStore := pg.NewWorkoutStore(tx)
			for _, workout := range seed.Workouts {
				if _, err := workoutStore.Create(workout); err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			log.Println(err)
		}
	},
}
