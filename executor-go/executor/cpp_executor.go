package executor

import "time"

type CppExecutor struct {
	BaseExecutor
}

func (c CppExecutor) Execute(code string) (string, error) {
	return c.runLocalCommand(
		code,
		"main.cpp",
		[]string{"sh", "-c", "g++ {file} -o {dir}/a.out && timeout 3s {dir}/a.out"},
		8*time.Second,
	)
}
