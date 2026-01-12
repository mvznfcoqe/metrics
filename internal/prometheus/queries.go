package prometheus

import (
	"context"
	"fmt"
	"time"

	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
)

const NetworkRateInterval = "20s"
const NetworkDeviceFilter = `lo|docker.*|taislcale.*|veth.*|br-.*`
const NodeExportersJob = "node-exporters"

const (
	NodesQuery                   = `up{job="` + NodeExportersJob + `"}`
	NodeUnameQuery               = `node_uname_info{job="` + NodeExportersJob + `"} or last_over_time(node_uname_info{job="` + NodeExportersJob + `"}[4d])`
	UptimeQueryTemplate          = `time() - node_boot_time_seconds{instance="%s"}`
	NetworkDownloadQueryTemplate = `rate(node_network_receive_bytes_total{instance="%s",device!~"%s"}[%s])`
	NetworkUploadQueryTemplate   = `rate(node_network_transmit_bytes_total{instance="%s",device!~"%s"}[%s])`
	ContainersQuery              = `time() - container_start_time_seconds{instance="%s",name!=""}`
)

func QueryVector(ctx context.Context, api v1.API, query string) (model.Vector, error) {
	result, _, err := api.Query(ctx, query, time.Now())

	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}

	vector, ok := result.(model.Vector)

	if !ok {
		return nil, fmt.Errorf("unexpected result type: expected Vector, got %T", result)
	}

	return vector, nil
}

func GetNodeInfo(ctx context.Context, api v1.API) (model.Vector, error) {
	return QueryVector(ctx, api, NodesQuery)
}

func GetNodeUname(ctx context.Context, api v1.API) (model.Vector, error) {
	return QueryVector(ctx, api, NodeUnameQuery)
}

func GetNodeUptime(ctx context.Context, api v1.API, instance string) (model.Vector, error) {
	query := fmt.Sprintf(UptimeQueryTemplate, instance)

	return QueryVector(ctx, api, query)
}

func GetNetworkDownloadRate(ctx context.Context, api v1.API, instance string) (model.Vector, error) {
	query := fmt.Sprintf(NetworkDownloadQueryTemplate, instance, NetworkDeviceFilter, NetworkRateInterval)

	return QueryVector(ctx, api, query)
}

func GetNetworkUploadRate(ctx context.Context, api v1.API, instance string) (model.Vector, error) {
	query := fmt.Sprintf(NetworkUploadQueryTemplate, instance, NetworkDeviceFilter, NetworkRateInterval)

	return QueryVector(ctx, api, query)
}

func GetContainers(ctx context.Context, api v1.API, instance string) (model.Vector, error) {
	query := fmt.Sprintf(ContainersQuery, instance)

	return QueryVector(ctx, api, query)
}
