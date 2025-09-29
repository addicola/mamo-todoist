package repositories

import "backend/mamo/entities"

type TodoDARepository struct{}

func NewTodoDARepository() TodoRepository {
	return &TodoDARepository{}
}

func (r *TodoDARepository) List() ([]*entities.Todo, error) {
	todos := []*entities.Todo{
		{Title: "Learn Go"},
		{Title: "Build a web app"},
	}
	return todos, nil
}
