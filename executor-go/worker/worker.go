package worker

import (
	"executor/factory"
	"executor/model"
	"strings"
	"time"
)

func StartWorker(jobs chan model.Job) {
	for job := range jobs {

		executor := factory.GetExecutor(job.Language)

		if executor == nil {
			job.Result <- model.Result{
				Output: "Unsupported language",
				Status: "error",
				Time:   "N/A",
			}
			continue
		}
		start := time.Now()
		output, err := executor.Execute(job.Code)
		status := "success"

		if err != nil {
			errStr := err.Error()
			outStr := string(output)

			switch {
			case strings.Contains(errStr, "timeout") ||
				strings.Contains(errStr, "context deadline exceeded"):
				status = "timeout"

			case strings.Contains(outStr, "error:") ||
				strings.Contains(outStr, "SyntaxError") ||
				strings.Contains(outStr, "compilation failed"):
				status = "compile_error"

			case strings.Contains(outStr, "Exception") ||
				strings.Contains(outStr, "Traceback") ||
				strings.Contains(outStr, "panic") ||
				strings.Contains(outStr, "runtime error"):
				status = "runtime_error"

			default:
				status = "runtime_error"
			}
		}
		execTime := time.Since(start).String()
		job.Result <- model.Result{
			Output: string(output),
			Status: status,
			Time:   execTime,
		}
	}
}
