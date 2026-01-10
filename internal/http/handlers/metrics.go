package handlers

import (
	"encoding/json"
	"metrics/internal/models"
	"metrics/internal/service"
	"net/http"
)

type MetricsHandler struct {
	metrics *service.MetricsService
}

func NewMetricsHandler(m *service.MetricsService) *MetricsHandler {
	return &MetricsHandler{metrics: m}
}

type GetAllNodesResponse struct {
	Nodes []models.Node `json:"nodes,omitempty"`
}

// @Summary Get all nodes
// @Description Returns array of nodes with metrics
// @Tags metrics
// @Produce json
// @Success 200 {object} GetAllNodesResponse
// @Failure 500 {object} object
// @Router /metrics/nodes [get]
func (h *MetricsHandler) GetAllNodes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	nodes, err := h.metrics.GetAllNodes(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(struct{}{})
		return
	}

	json.NewEncoder(w).Encode(GetAllNodesResponse{Nodes: nodes})
}
