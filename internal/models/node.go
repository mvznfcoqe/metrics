package models

type NodeNetwork struct {
	Download int `json:"download" binding:"required"`
	Upload   int `json:"upload" binding:"required"`
}

type Node struct {
	Location   string      `json:"location" binding:"required"`
	Provider   string      `json:"provider" binding:"required"`
	Uptime     float64     `json:"uptime" binding:"required"`
	Network    NodeNetwork `json:"network" binding:"required"`
	Containers []Container `json:"containers" binding:"required"`
}
