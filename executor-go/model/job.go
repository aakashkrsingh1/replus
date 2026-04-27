package model

type Job struct {
	Code     string
	Language string
	Result   chan Result
}

type Result struct {
	Output string
	Status string
	Time   string
}
