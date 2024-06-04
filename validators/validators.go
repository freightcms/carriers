package validators

import "context"

type ValidationErrors interface {
	// JSON Creates a new JSON interface which can then be appended to the responses of certain objects or be output tot he user.
	JSON(ctx context.Context) interface{}
}

// Validator is the interface that can be implemented for schemas or models during initial call for creating, deleting, or updating.
// Validators can expand a wide range of use cases, form user permissions with passed database context or user principles
type ValidatorFunc func(ctx context.Context, val any) []ValidationErrors
