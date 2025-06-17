package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"github.com/vegitobluefan/task-manager/domain"
	"github.com/vegitobluefan/task-manager/usecase"
)

func SetupRoutes(r *gin.Engine, uc usecase.TaskUseCase, repo domain.TaskRepository) {
	r.POST("/tasks", func(c *gin.Context) {
		var input struct {
			Type    string `json:"type"`
			Payload string `json:"payload"`
		}

		if err := c.BindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		task := &domain.Task{
			ID:      uuid.New().String(),
			Type:    input.Type,
			Payload: input.Payload,
			Status:  "queued",
		}

		id, err := uc.Enqueue(task)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusAccepted, gin.H{"id": id})
	})

	r.GET("/tasks/:id", func(c *gin.Context) {
		id := c.Param("id")
		task, err := repo.GetByID(id)
		if err != nil || task == nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "task not found"})
			return
		}
		c.JSON(http.StatusOK, task)
	})

	r.GET("/tasks", func(c *gin.Context) {
		tasks, err := repo.ListTasks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "could not retrieve tasks"})
			return
		}
		c.JSON(http.StatusOK, tasks)
	})
}
