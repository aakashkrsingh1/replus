package executor

import "time"

type PythonExecutor struct {
	BaseExecutor
}

func (p PythonExecutor) Execute(code string) (string, error) {
	return p.runLocalCommand(
		code,
		"main.py",
		[]string{"python3", "{file}"},
		2*time.Second,
	)
}
