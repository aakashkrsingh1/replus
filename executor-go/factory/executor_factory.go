package factory

import "executor/executor"

func GetExecutor(language string) executor.Executor {
	switch language {
	case "python":
		return executor.PythonExecutor{}
	case "java":
		return executor.JavaExecutor{}
	case "cpp":
		return executor.CppExecutor{}
	default:
		return nil
	}
}
