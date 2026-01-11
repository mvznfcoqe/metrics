package models

type Container struct {
	Name    string `json:"name" binding:"required"`
	Project string `json:"project" binding:"required"`
}
