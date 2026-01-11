package handlers

import (
	"encoding/json"
	"metrics/internal/models"
	"net/http"
)

type ProvidersHandler struct {
	providers map[string]models.Provider
}

func NewProvidersHandler() *ProvidersHandler {
	providers := map[string]models.Provider{
		"cloud": {
			Link: "https://cloud.ru/",
		},
		"veesp": {
			Link: "https://veesp.com/",
		},
		"nexus": {
			Link: "https://h2.nexus/",
		},
		"llhost": {
			Link: "https://llhost.eu/",
		},
		"hostingrussia": {
			Link: "https://hosting-russia.ru/",
		},
	}

	return &ProvidersHandler{providers: providers}
}

// @Summary Get all providers
// @Description Returns object with provider names as keys and provider information as values
// @Tags providers
// @Produce json
// @Success 200 {object} map[string]models.Provider
// @Router /providers [get]
func (h *ProvidersHandler) GetProviders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.providers)
}
