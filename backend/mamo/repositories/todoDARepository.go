package repositories

import "backend/mamo/entities"

type TodoCelestiaRepository struct{}

func NewTodoCelestiaRepository() TodoRepository {
	return &TodoCelestiaRepository{}
}

func (r *TodoCelestiaRepository) List() ([]*entities.Todo, error) {
	todos := []*entities.Todo{
		{Title: "Learn Go"},
		{Title: "Build a web app"},
	}
	return todos, nil
}
