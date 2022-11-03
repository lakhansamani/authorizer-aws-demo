package db

import (
	"apis/models"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

const (
	todoTableName = "todos"
)

// AddTodo database helper to add todo
func AddTodo(ctx context.Context, todo *models.Todo) (*models.Todo, error) {
	if todo.ID == "" {
		todo.ID = uuid.New().String()
	}
	todo.CreatedAt = time.Now().Unix()
	todo.UpdatedAt = time.Now().Unix()
	todo.IsCompleted = false

	table := provider.Table(todoTableName)

	err := table.Put(todo).RunWithContext(ctx)

	if err != nil {
		return nil, err
	}
	return todo, nil
}

// UpdateTodo database helper to update todo
func UpdateTodo(ctx context.Context, todo *models.Todo) (*models.Todo, error) {
	table := provider.Table(todoTableName)

	todo.UpdatedAt = time.Now().Unix()

	err := UpdateByHashKey(table, "id", todo.ID, todo)
	if err != nil {
		return todo, err
	}

	if err != nil {
		return todo, err
	}
	return todo, nil
}

// GetTodos database helper to get list of todo as per userID
func GetTodos(ctx context.Context, userID string) ([]*models.Todo, error) {
	table := provider.Table(todoTableName)
	todos := []*models.Todo{}
	var todo models.Todo

	iter := table.Scan().Index("user_id").Filter("'user_id' = ?", userID).Iter()
	err := iter.Err()
	if err != nil {
		return nil, err
	}

	for iter.NextWithContext(ctx, &todo) {
		todoApi := &models.Todo{
			ID:          todo.ID,
			Title:       todo.Title,
			IsCompleted: todo.IsCompleted,
			CreatedAt:   todo.CreatedAt,
			UpdatedAt:   todo.UpdatedAt,
		}
		todos = append(todos, todoApi)
	}

	return todos, nil
}

// DeleteTodo database helper to delete todo as per given todoID
func DeleteTodo(ctx context.Context, todoID string) error {
	table := provider.Table(todoTableName)
	err := table.Delete("id", todoID).Run()
	if err != nil {
		return err
	}
	return nil
}

func GetTodoById(ctx context.Context, todoID string) (*models.Todo, error) {
	table := provider.Table(todoTableName)
	var todo *models.Todo

	err := table.Get("id", todoID).OneWithContext(ctx, &todo)
	if err != nil {
		return nil, err
	}

	if todo.ID == "" {
		return nil, errors.New("todo not found")
	}

	return todo, nil
}
