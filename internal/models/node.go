package models

type NodeNetwork struct {
	Download int `json:"download"`
	Upload   int `json:"upload"`
}

type Node struct {
	Location   string      `json:"location"`
	Provider   string      `json:"provider"`
	Uptime     float64     `json:"uptime"`
	Network    NodeNetwork `json:"network"`
	Containers []Container `json:"containers"`
}
