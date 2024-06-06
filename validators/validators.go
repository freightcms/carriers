package validators

import "context"

// Validator is the interface that can be implemented for schemas or models during initial call for creating, deleting, or updating.
// Validators can expand a wide range of use cases, form user permissions with passed database context or user principles
type ValidatorFunc func(ctx context.Context, val any) []error

func CreateValidatorFunc(callable func(ctx context.Context, val any) []error) ValidatorFunc {
	return ValidatorFunc(callable)
}
