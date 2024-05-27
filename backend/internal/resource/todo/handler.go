package todo

import (
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	common "github.com/charukak/todo-app-htmx/common/pkg"
	"github.com/gin-gonic/gin"
)

type ITodoService interface {
	GetAll() ([]common.Todo, error)
	GetByID(id int) (common.Todo, error)
	Create(todo common.Todo) (common.Todo, error)
	Update(todo common.Todo) (common.Todo, error)
	Delete(id int) error
}

type TodoService struct {
	db *sql.DB
}

func (s *TodoService) GetAll() ([]common.Todo, error) {
	fmt.Println("Calling GET /todos")
	rows, err := s.db.Query("SELECT * FROM todos")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []common.Todo = []common.Todo{}

	for rows.Next() {
		var todo common.Todo

		err := rows.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status)

		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	return todos, nil

}

func (s *TodoService) GetByID(id int) (common.Todo, error) {
	fmt.Printf("Calling GET /todos/%d\n", id)
	fmt.Println("New log")
	var todo common.Todo

	row := s.db.QueryRow("SELECT * FROM todos WHERE id = ?", id)

	err := row.Scan(&todo.ID, &todo.Title, &todo.Description, &todo.Status)

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (s *TodoService) Create(todo common.Todo) (common.Todo, error) {
	fmt.Printf("Calling Post /todos title=%s, description=%s\n",todo.Title, todo.Description)

	result, err := s.db.Exec(
		"INSERT INTO todos (title, description, status) VALUES (?, ?, ?)",
		todo.Title,
		todo.Description,
		todo.Status)

	if err != nil {
		return todo, err
	}

	lastID, err := result.LastInsertId()

	if err != nil {
		return todo, err
	}

	todo.ID = int(lastID)

	return todo, nil
}

func (s *TodoService) Update(todo common.Todo) (common.Todo, error) {
	cols := []string{}

	if todo.Title != "" {
		cols = append(cols, "title = '"+todo.Title+"'")
	}

	if todo.Description != "" {
		cols = append(cols, "description = '"+todo.Description+"'")
	}

	cols = append(cols, "status = "+strconv.FormatBool(todo.Status))

	_, err := s.db.Exec(
		fmt.Sprintf("UPDATE todos SET %s WHERE id = ?", strings.Join(cols, ",")),
		todo.ID)

	if err != nil {
		return todo, err
	}

	return todo, nil
}

func (s *TodoService) Delete(id int) error {
	fmt.Sprintf("Calling Delete /todos/%d\n",id)
	_, err := s.db.Exec("DELETE FROM todos WHERE id = ?", id)

	if err != nil {
		return err
	}

	return nil
}

func NewTodoService(db *sql.DB) ITodoService {
	return &TodoService{db}
}

func RegisterTodoHandler(r *gin.Engine, db *sql.DB) {
	service := NewTodoService(db)

	r.GET("/todos", func(c *gin.Context) {
		todos, err := service.GetAll()

		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		c.JSON(200, todos)
	})

	r.GET("/todos/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}

		todo, err := service.GetByID(id)

		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		c.JSON(200, todo)
	})

	r.POST("/todos", func(c *gin.Context) {
		var todo common.Todo

		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}

		newTodo, err := service.Create(todo)

		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		c.JSON(201, newTodo)
	})

	r.PUT("/todos/:id", func(c *gin.Context) {
		var todo common.Todo

		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}

		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}

		todo.ID = id

		newTodo, err := service.Update(todo)

		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		c.JSON(200, newTodo)
	})

	r.DELETE("/todos/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))

		if err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}

		err = service.Delete(id)

		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		c.JSON(200, gin.H{"message": "common.Todo deleted"})
	})
}
