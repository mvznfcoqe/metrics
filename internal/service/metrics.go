package service

import (
	"context"
	"fmt"
	"log"
	"math"
	"metrics/internal/config"
	"metrics/internal/models"
	"metrics/internal/prometheus"
	"strings"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
)

type MetricsService struct {
	prometheusAPI v1.API
}

func NewMetricsService(prometheusAPI v1.API) *MetricsService {
	return &MetricsService{
		prometheusAPI,
	}
}

func (s *MetricsService) parseNodename(nodename string) (models.Node, error) {
	parts := strings.Split(nodename, "-")

	if len(parts) < 2 {
		return models.Node{}, fmt.Errorf("invalid nodename format: %s", nodename)
	}

	return models.Node{
		Provider: parts[0],
		Location: parts[1],
	}, nil
}

func (s *MetricsService) getNodeUptime(ctx context.Context, instance string) (float64, error) {
	uptimeResult, error := prometheus.GetNodeUptime(ctx, s.prometheusAPI, instance)

	if error != nil {
		return 0, error
	}

	uptime := math.Floor(float64(uptimeResult[0].Value))

	return uptime, nil
}

func (s *MetricsService) getNodeNetworkStatistics(ctx context.Context, instance string) (models.NodeNetwork, error) {
	downloadResult, error := prometheus.GetNetworkDownloadRate(ctx, s.prometheusAPI, instance)

	if error != nil {
		return models.NodeNetwork{}, error
	}

	uploadResult, error := prometheus.GetNetworkUploadRate(ctx, s.prometheusAPI, instance)

	if error != nil {
		return models.NodeNetwork{}, error
	}

	download := int(math.Floor(float64(downloadResult[0].Value)))
	upload := int(math.Floor(float64(uploadResult[0].Value)))

	return models.NodeNetwork{
		Download: download,
		Upload:   upload,
	}, nil
}

func (s *MetricsService) getNodeContainers(ctx context.Context, instance string) ([]models.Container, error) {
	containersResult, error := prometheus.GetContainers(ctx, s.prometheusAPI, instance)

	if error != nil {
		return nil, error
	}

	containers := make([]models.Container, 0, len(containersResult))

	for _, sample := range containersResult {
		name := string(sample.Metric["name"])
		project := string(sample.Metric["container_label_project"])

		containers = append(containers, models.Container{Name: name, Project: project})
	}

	return containers, nil
}

func (s *MetricsService) GetAllNodes(ctx context.Context) ([]models.Node, error) {
	config, error := config.Load()

	if error != nil {
		log.Fatal("Failed to get configuration")

		return nil, error
	}

	nodenameResult, error := prometheus.GetNodeInfo(ctx, s.prometheusAPI)

	if error != nil {
		return nil, error
	}

	nodes := make([]models.Node, 0, len(nodenameResult))

	for _, sample := range nodenameResult {
		nodename := string(sample.Metric["nodename"])
		nodeExporterInstance := string(sample.Metric["instance"])
		cadvisorInstance := strings.Replace(nodeExporterInstance, config.NodeExporterPort, config.CadvisorPort, 1)

		node, error := s.parseNodename(nodename)
		if error != nil {
			log.Printf("Warning: skipping node %s: %v", nodename, error)
			continue
		}

		uptime, error := s.getNodeUptime(ctx, nodeExporterInstance)
		if error != nil {
			log.Printf("Warning: failed to get uptime for %s: %v", nodeExporterInstance, error)
			continue
		}
		node.Uptime = uptime

		network, error := s.getNodeNetworkStatistics(ctx, nodeExporterInstance)
		if error != nil {
			log.Printf("Warning: failed to get network stats for %s: %v", nodeExporterInstance, error)
			continue
		}
		node.Network = network

		containers, error := s.getNodeContainers(ctx, cadvisorInstance)
		if error != nil {
			log.Printf("Warning: failed to get containers for %s: %v", cadvisorInstance, error)
		}
		node.Containers = containers

		nodes = append(nodes, node)
	}

	return nodes, nil
}
