package model

// MovementStore defines an interface for a store of Movements
type MovementStore interface {
	Create(movement *Movement) (string, error)
	Get(id string) (*Movement, error)
	List() ([]Movement, error)
	Update(movement *Movement) error
	Delete(id string) error
}
