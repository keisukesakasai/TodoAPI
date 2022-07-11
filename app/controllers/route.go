package controllers

import (
	"log"
	"net/http"
	"todoapi/app/models"

	"github.com/gin-gonic/gin"
)

type createTodoRequest struct {
	Content string `json:"content"`
	User_Id string `json:"user_id"`
}

type getTodosByUserRequest struct {
	User_Id string `json:"user_id"`
}

type getTodoRequest struct {
	Todo_Id string `json:"todo_id"`
}

type updateTodoRequest struct {
	Content string `json:"content"`
	User_Id string `json:"user_id"`
	Todo_Id string `json:"todo_id"`
}

type deleteTodoRequest struct {
	Todo_Id string `json:"todo_id`
}

func createTodo(c *gin.Context) {
	_, span := tracer.Start(c.Request.Context(), "TODO 登録")
	defer span.End()

	var createTodorequest createTodoRequest
	if err := c.BindJSON(&createTodorequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("---createTodo---")
	log.Println(createTodorequest)

	content := createTodorequest.Content
	user_id := createTodorequest.User_Id
	if err := models.CreateTodo(c, content, user_id); err != nil {
		log.Println(err)
	}
	log.Println("TODO 登録")

	c.JSON(http.StatusOK, gin.H{
		"content": content,
	})
}

func updateTodo(c *gin.Context) {
	_, span := tracer.Start(c.Request.Context(), "TODO 更新")
	defer span.End()

	var updateTodorequest updateTodoRequest
	if err := c.BindJSON(&updateTodorequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	content := updateTodorequest.Content
	user_id := updateTodorequest.User_Id
	todo_id := updateTodorequest.Todo_Id

	if err := models.UpdateTodo(c, content, user_id, todo_id); err != nil {
		log.Println(err)
	}
	log.Println("TODO 登録")

	c.JSON(http.StatusOK, gin.H{
		"content": content,
	})
}

func getTodos(c *gin.Context) {

}

func getTodo(c *gin.Context) {
	_, span := tracer.Start(c.Request.Context(), "TODO 参照")
	defer span.End()

	var getTodorequest getTodoRequest
	if err := c.BindJSON(&getTodorequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(getTodorequest)
	todo_id := getTodorequest.Todo_Id

	todo, err := models.GetTodo(c, todo_id)
	if err != nil {
		log.Println(err)
	}
	log.Println("TODO 参照")

	log.Println(todo)

	c.JSON(http.StatusOK, gin.H{
		"ID":        todo.ID,
		"Content":   todo.Content,
		"UserID":    todo.UserID,
		"CreatedAt": todo.CreatedAt,
	})
}

func deleteTodo(c *gin.Context) {
	_, span := tracer.Start(c.Request.Context(), "TODO 削除")
	defer span.End()

	var deleteTodorequest deleteTodoRequest
	if err := c.BindJSON(&deleteTodorequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(deleteTodorequest)
	todo_id := deleteTodorequest.Todo_Id

	todo, err := models.GetTodo(c, todo_id)
	if err != nil {
		log.Println(err)
	}
	log.Println("TODO 参照")

	if todo.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"resultCode": "Todoがありません",
		})
	} else {
		err := models.DeleteTodo(c, todo_id)
		if err != nil {
			log.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"resultCode": "ID : " + todo_id + " の Todo を正常に削除しました",
		})
	}
}

func getTodosByUser(c *gin.Context) {
	_, span := tracer.Start(c.Request.Context(), "ユーザごとの TODO 参照")
	defer span.End()

	var getTodosByUserrequest getTodosByUserRequest
	if err := c.BindJSON(&getTodosByUserrequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println(getTodosByUserrequest)
	user_id := getTodosByUserrequest.User_Id

	todos, err := models.GetTodosByUser(c, user_id)
	if err != nil {
		log.Println(err)
	}
	log.Println("ユーザごとの TODO 参照")

	log.Println(todos)

	c.JSON(http.StatusOK, gin.H{
		"todos": todos,
	})
}
