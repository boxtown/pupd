package pg

import (
	"testing"

	"github.com/boxtown/pupd/model"
)

func TestBuildCreateExerciseQuery(t *testing.T) {
	exerciseIDs := []string{"1", "2"}
	exercises := []*model.Exercise{
		&model.Exercise{
			Movement: &model.Movement{ID: "1"},
		},
		&model.Exercise{
			Movement: &model.Movement{ID: "2"},
		},
	}
	query, args := buildCreateExercisesQuery("1", exerciseIDs, exercises)
	expectedQuery := "INSERT INTO core.exercises (exercise_id, workout_id, pos, movement_id)" +
		" VALUES ($1, $2, $3, $4), ($5, $6, $7, $8)"
	if query != expectedQuery {
		t.Errorf("Expected %s, got %s", expectedQuery, query)
	}
	expectedArgs := []interface{}{
		"1", "1", 0, "1",
		"2", "1", 1, "2",
	}
	for i, arg := range args {
		if arg != expectedArgs[i] {
			t.Errorf("Expected arg %s at index %d, got %s", expectedArgs[i], i, arg)
		}
	}
}

func TestBuildCreateExerciseSetsQuery(t *testing.T) {
	exerciseIDs := []string{"1", "2"}
	exercises := []*model.Exercise{
		&model.Exercise{
			Sets: []*model.ExerciseSet{
				&model.ExerciseSet{
					Reps: 5,
				},
				&model.ExerciseSet{
					Reps: 3,
				},
			},
		},
		&model.Exercise{
			Sets: []*model.ExerciseSet{
				&model.ExerciseSet{
					Reps: 3,
				},
			},
		},
	}
	query, args := buildCreateExerciseSetsQuery(exerciseIDs, exercises)
	if err != nil {
		t.Fatal(err.Error())
	}
	expectedQuery := "INSERT INTO core.exercises_sets (exercise_id, pos, reps)" +
		" VALUES ($1, $2, $3), ($4, $5, $6), ($7, $8, $9)"
	if query != expectedQuery {
		t.Errorf("Expected %s, got %s", expectedQuery, query)
	}
	expectedArgs := []interface{}{
		"1", 0, 5,
		"1", 1, 3,
		"2", 0, 3,
	}
	for i, arg := range args {
		if arg != expectedArgs[i] {
			t.Errorf("Expected arg %s at index %d, got %s", expectedArgs[i], i, arg)
		}
	}
}
