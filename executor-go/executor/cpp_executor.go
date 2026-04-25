package executor

type CppExecutor struct {
	BaseExecutor
}

func (c CppExecutor) Execute(code string) (string, error) {
	return c.runLocalCommand(
		code,
		"main.cpp",
		[]string{"sh", "-c", "g++ {file} -o {dir}/a.out && {dir}/a.out"},
	)
}
