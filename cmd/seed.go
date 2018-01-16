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
		movements := []*model.Movement{
			&model.Movement{Name: "Back Squat"},
			&model.Movement{Name: "Deadlift"},
			&model.Movement{Name: "Bench Press"},
			&model.Movement{Name: "Overhead Press"},
		}
		units := []*model.Unit{
			&model.Unit{Name: "%"},
		}

		url := fmt.Sprintf("postgres://%s:%s@localhost/pupd", seedUser, seedPass)
		source, err := pg.NewDataSource(url)
		if err != nil {
			log.Fatal(err.Error())
		}
		defer source.Close()

		err = source.Transaction(func(tx *sqlx.Tx) error {
			movementStore := pg.NewMovementStore(tx)
			for _, movement := range movements {
				id, err := movementStore.Create(movement)
				if err != nil {
					return err
				}
				movement.ID = id
			}
			unitStore := pg.NewUnitStore(tx)
			for _, unit := range units {
				id, err := unitStore.Create(unit)
				if err != nil {
					return err
				}
				unit.ID = id
			}

			// Workouts is created mid-transaction in order to make
			// use of generated IDs
			workouts := []*model.Workout{
				&model.Workout{
					Name: "Sample Workout",
					Exercises: []*model.Exercise{
						&model.Exercise{
							Pos:      0,
							Movement: movements[0],
							Sets: []*model.ExerciseSet{
								&model.ExerciseSet{
									Pos:          0,
									Reps:         5,
									MinIntensity: 0.9,
									Unit:         units[0],
								},
							},
						},
					},
				},
			}

			workoutStore := pg.NewWorkoutStore(tx)
			for _, workout := range workouts {
				if _, err := workoutStore.Create(workout); err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			log.Println(err.Error())
		}
	},
}
