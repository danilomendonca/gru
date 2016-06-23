package analyzer

import (
	"errors"

	log "github.com/elleFlorio/gru/Godeps/_workspace/src/github.com/Sirupsen/logrus"

	evl "github.com/elleFlorio/gru/autonomic/analyzer/evaluator"
	"github.com/elleFlorio/gru/data"
	srv "github.com/elleFlorio/gru/service"
)

var (
	analytics             data.GruAnalytics
	ErrNoRunningInstances error = errors.New("No active instance to analyze")
)

func init() {
	analytics = data.GruAnalytics{}
}

// TODO - What if I have no stats?
func Run(stats data.GruStats) data.Shared {
	log.WithField("status", "init").Debugln("Gru Monitor")
	defer log.WithField("status", "done").Debugln("Gru Monitor")

	if len(stats.Metrics.Service) == 0 {
		log.Debugln("No stats to compute")
		return data.Shared{}
	}

	analytics.Service = computeServicesAnalytics(stats.Metrics.Service)
	analytics.System = computeSystemAnalytics(stats.Metrics.System)
	shared := computeSharedData(analytics)

	return shared
}

func computeServicesAnalytics(servStats map[string]data.MetricData) map[string]data.AnalyticData {
	servAnalytics := make(map[string]data.AnalyticData)

	for service, metrics := range servStats {
		aData := data.AnalyticData{}
		baseAnalytics := metrics.BaseMetrics
		userAnalytics := evl.ComputeMetricAnalytics(service, metrics.UserMetrics)
		aData.BaseAnalytics = baseAnalytics
		aData.UserAnalytics = userAnalytics

		servAnalytics[service] = aData
	}

	return servAnalytics
}

func computeSystemAnalytics(sysStats data.MetricData) data.AnalyticData {
	sysAnalitycs := data.AnalyticData{
		BaseAnalytics: sysStats.BaseMetrics,
		UserAnalytics: sysStats.UserMetrics,
	}

	return sysAnalitycs
}

func computeSharedData(analytics data.GruAnalytics) data.Shared {
	local := computeLocaShared(analytics)
	cluster := computeClusterShared(local)

	return cluster
}

func computeLocaShared(analytics data.GruAnalytics) data.Shared {
	local := data.Shared{}
	srvActive := []string{}
	for name, values := range analytics.Service {
		srvShared := data.ServiceShared{}
		srvShared.Data.BaseShared = values.BaseAnalytics
		srvShared.Data.UserShared = values.UserAnalytics
		srvShared.Active = srv.IsServiceActive(name)

		local.Service[name] = srvShared

		if srvShared.Active {
			srvActive = append(srvActive, name)
		}
	}

	local.System.Data.BaseShared = analytics.System.BaseAnalytics
	local.System.Data.UserShared = analytics.System.UserAnalytics
	local.System.ActiveServices = srvActive

	data.SaveSharedLocal(local)

	return local
}

func computeClusterShared(local data.Shared) data.Shared {
	storedCluster, err := data.GetSharedCluster()
	if err != nil {
		log.WithField("err", err).Debugln("Cannot compute cluster data")
		return local
	}

	toMerge := []data.Shared{local, storedCluster}
	cluster, err := data.MergeShared(toMerge)
	if err != nil {
		return local
	}

	data.SaveSharedCluster(cluster)

	return cluster
}
