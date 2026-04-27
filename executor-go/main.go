package main

import (
	"net/http"
	"time"

	"executor/model"
	"executor/worker"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Code     string `json:"code"`
	Language string `json:"language"`
}

func main() {

	jobs := make(chan model.Job, 100)

	for i := 0; i < 5; i++ {
		go worker.StartWorker(jobs)
	}

	r := gin.Default()

	r.POST("/execute", func(c *gin.Context) {
		var req Request

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		resultChan := make(chan model.Result)
		start := time.Now()

		job := model.Job{
			Code:     req.Code,
			Language: req.Language,
			Result:   resultChan,
		}

		jobs <- job

		res := <-resultChan
		duration := time.Since(start)

		c.JSON(http.StatusOK, gin.H{
			"output":         res.Output,
			"status":         res.Status,
			"execution_time": res.Time,
			"total_time":     duration.String(),
		})
	})

	r.Run(":8081")
}
