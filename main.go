package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Name: "Take a shower", Completed: true},
	{ID: "2", Name: "Cook dinner", Completed: false},
}

func GetToDos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, todos)
	// context.Writer.Write([]byte("hello"))
}

func GetToDo(context *gin.Context) {
	id := context.Param("id")
	needToDo, err := findToDo(id)
	if err != nil {
		return
	}
	context.IndentedJSON(http.StatusOK, needToDo)
}
func findToDo(id string) (*todo, error) {
	for _, j := range todos {
		if j.ID == id {
			return &j, nil
		}
	}
	return nil, errors.New("")
}

func PostToDo(context *gin.Context) {
	var newTodo todo
	if err := context.BindJSON(&newTodo); err != nil {
		return
	}
	todos = append(todos, newTodo)
	context.IndentedJSON(http.StatusCreated, newTodo)
}

func main() {
	router := gin.Default()
	router.GET("/todos", GetToDos)
	router.GET("todos/:id", GetToDo)
	router.POST("/new", PostToDo)
	log.Fatalln(router.Run(":80"))
}
