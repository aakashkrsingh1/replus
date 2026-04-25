package executor

type Executor interface {
	Execute(code string) (string, error)
}
