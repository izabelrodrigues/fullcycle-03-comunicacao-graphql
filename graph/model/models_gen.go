// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

// Inputs
type NewCategory struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type NewChapter struct {
	Name     string `json:"name"`
	CourseID string `json:"courseId"`
}

type NewCourse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	CategoryID  string `json:"categoryId"`
}
