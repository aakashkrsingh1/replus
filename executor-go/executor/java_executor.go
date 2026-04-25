package executor

type JavaExecutor struct {
	BaseExecutor
}

func (j JavaExecutor) Execute(code string) (string, error) {
	return j.runLocalCommand(
		code,
		"Main.java",
		[]string{"sh", "-c", "javac {file} && java -cp {dir} Main"},
	)
}
