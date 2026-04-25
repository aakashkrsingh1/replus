package worker

import (
	"executor/factory"
	"executor/model"
)

func StartWorker(jobs chan model.Job) {
	for job := range jobs {

		executor := factory.GetExecutor(job.Language)

		if executor == nil {
			job.Result <- "Unsupported language"
			continue
		}

		output, _ := executor.Execute(job.Code)

		job.Result <- output
	}
}
