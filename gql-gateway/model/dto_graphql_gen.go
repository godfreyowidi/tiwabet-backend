// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Mutation struct {
}

type NewUser struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Query struct {
}

type UpdateUser struct {
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
}
