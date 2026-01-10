package main

import (
	"log"
	"metrics/internal/config"
	"metrics/internal/http/handlers"
	"metrics/internal/prometheus"
	"metrics/internal/service"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	config, err := config.Load()

	if err != nil {
		log.Fatal("Failed to get configuration")

		return
	}

	prometheusClient, err := prometheus.NewClient(config.PrometheusURL)

	if err != nil {
		log.Fatalf("Failed to initalize Prometheus client: %v", err)

		return
	}

	metricsService := service.NewMetricsService(prometheusClient.API())
	metricsHandler := handlers.NewMetricsHandler(metricsService)

	mux := http.NewServeMux()
	mux.HandleFunc("/metrics/nodes", metricsHandler.GetAllNodes)

	c := cors.New(cors.Options{
		AllowedOrigins:   config.AllowedOrigins,
		AllowedMethods:   []string{"GET"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	log.Printf(`Server starting on: %s`, config.Port)
	log.Fatal(http.ListenAndServe(":"+config.Port, c.Handler(mux)))
}
