package repositories

import "backend/mamo/entities"

type TodoRepository interface {
	List() ([]*entities.Todo, error)
}
