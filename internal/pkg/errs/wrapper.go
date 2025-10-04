package errs

import (
	"fmt"
	"strings"
)

// unionErrors combines target error with additional errors using format rule (%w/%v) for flexible composition
func unionErrors(target error, rule string, additional ...error) error {
	fStr := "%w" + strings.Repeat(rule, len(additional))
	chain := []any{target}
	for _, err := range additional {
		chain = append(chain, err)
	}
	return fmt.Errorf(fStr, chain...)
}

// WrapErrors creates error chain using %w wrapping. Format: "target: err1: ..."
func WrapErrors(target error, additional ...error) error {
	return unionErrors(target, ": %w", additional...)
}

// ConcatErrors concatenates errors with %v formatting. Format: "target: err1: ..."
func ConcatErrors(target error, additional ...error) error {
	return unionErrors(target, ": %v", additional...)
}
