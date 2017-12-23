package model

type MovementStore interface {
	Create(movement *Movement) (string, error)
	Get(id string) (*Movement, error)
	List() ([]Movement, error)
	Update(movement *Movement) error
	Delete(id string) error
}
