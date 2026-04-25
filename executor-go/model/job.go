package model

type Job struct {
	Code     string
	Language string
	Result   chan string
}
