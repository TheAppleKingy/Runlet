package mocks

import "github.com/stretchr/testify/mock"

// GetMockArgs returns slice of mock.Anything
func GetMockArgs(count int) []any {
	args := make([]any, count)
	for i := range args {
		args[i] = mock.Anything
	}
	return args
}
