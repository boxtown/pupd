DROP SCHEMA IF EXISTS core CASCADE;
CREATE SCHEMA core;

SET search_path TO core;

CREATE TABLE IF NOT EXISTS users (
    user_id text PRIMARY KEY
);
GRANT ALL ON users TO pupd_api;

CREATE TABLE IF NOT EXISTS movements (
    movement_id uuid PRIMARY KEY,
    name text UNIQUE NOT NULL
);
GRANT ALL ON movements TO pupd_api;

CREATE TABLE IF NOT EXISTS workouts (
    workout_id uuid PRIMARY KEY,
    name text UNIQUE NOT NULL
);
GRANT ALL ON workouts TO pupd_api;

CREATE TABLE IF NOT EXISTS exercises (
    exercise_id uuid PRIMARY KEY,
    workout_id uuid NOT NULL REFERENCES workouts,
    pos integer NOT NULL,
    movement_id uuid NOT NULL REFERENCES movements,
    UNIQUE (workout_id, pos)
);
GRANT ALL ON exercises TO pupd_api;

CREATE TABLE IF NOT EXISTS exercise_sets (
    exercise_id uuid REFERENCES exercises,
    pos integer,
    reps integer NOT NULL,
    PRIMARY KEY (exercise_id, pos)
);
GRANT ALL ON exercise_sets TO pupd_api;