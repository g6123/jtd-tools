package util

type Result[T any] struct {
	Value T
	Err   error
}
