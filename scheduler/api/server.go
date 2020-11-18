package api

// these functions are the service side entry points for the thrift protocol
// (they are call from cloudscoot.go)

import (
	"github.com/apache/thrift/lib/go/thrift"
	log "github.com/sirupsen/logrus"

	"github.com/twitter/scoot/common/stats"
	"github.com/twitter/scoot/saga"
	schedthrift "github.com/twitter/scoot/scheduler/api/thrift"
	"github.com/twitter/scoot/scheduler/api/thrift/gen-go/scoot"
	"github.com/twitter/scoot/scheduler/server"
)

// NewHandler Creates and returns a new server Handler, which combines the scheduler,
// saga coordinator and stats receivers.
func NewHandler(scheduler server.Scheduler, sc saga.SagaCoordinator, stat stats.StatsReceiver) scoot.CloudScoot {
	handler := &Handler{scheduler: scheduler, sagaCoord: sc, stat: stat}
	go stats.StartUptimeReporting(stat, stats.SchedUptime_ms, stats.SchedServerStartedGauge, stats.DefaultStartupGaugeSpikeLen)
	return handler
}

// MakeServer Creates a Thrift server given a Handler and Thrift connection information
func MakeServer(handler scoot.CloudScoot,
	transport thrift.TServerTransport,
	transportFactory thrift.TTransportFactory,
	protocolFactory thrift.TProtocolFactory) thrift.TServer {
	return thrift.NewTSimpleServer4(
		scoot.NewCloudScootProcessor(handler),
		transport, transportFactory, protocolFactory)
}

// Handler Wrapping type that combines a scheduler, saga coordinator and stat receiver into a server
type Handler struct {
	scheduler server.Scheduler
	sagaCoord saga.SagaCoordinator
	stat      stats.StatsReceiver
}

// RunJob Implements RunJob Cloud Scoot API
func (h *Handler) RunJob(def *scoot.JobDefinition) (*scoot.JobId, error) {
	defer h.stat.Latency(stats.SchedServerRunJobLatency_ms).Time().Stop() // TODO errata metric - remove if unused
	h.stat.Counter(stats.SchedServerRunJobCounter).Inc(1)                 // TODO errata metric - remove if unused
	return schedthrift.RunJob(h.scheduler, def, h.stat)
}

// GetStatus Implements GetStatus Cloud Scoot API
func (h *Handler) GetStatus(jobID string) (*scoot.JobStatus, error) {
	defer h.stat.Latency(stats.SchedServerJobStatusLatency_ms).Time().Stop()
	h.stat.Counter(stats.SchedServerJobStatusCounter).Inc(1)
	return schedthrift.GetJobStatus(jobID, h.sagaCoord)
}

// KillJob Implements KillJob Cloud Scoot API
func (h *Handler) KillJob(jobID string) (*scoot.JobStatus, error) {
	defer h.stat.Latency(stats.SchedServerJobKillLatency_ms).Time().Stop()
	h.stat.Counter(stats.SchedServerJobKillCounter).Inc(1)
	return schedthrift.KillJob(jobID, h.scheduler, h.sagaCoord)
}

// OfflineWorker Implements OfflineWorker Cloud Scoot API
func (h *Handler) OfflineWorker(req *scoot.OfflineWorkerReq) error {
	return schedthrift.OfflineWorker(req, h.scheduler)
}

// ReinstateWorker Implements ReinstateWorker Cloud Scoot API
func (h *Handler) ReinstateWorker(req *scoot.ReinstateWorkerReq) error {
	return schedthrift.ReinstateWorker(req, h.scheduler)
}

// GetSchedulerStatus Implements GetSchedulerStatus Cloud Scoot API
func (h *Handler) GetSchedulerStatus() (*scoot.SchedulerStatus, error) {
	return schedthrift.GetSchedulerStatus(h.scheduler)
}

// SetSchedulerStatus  Implements SetSchedulerStatus Cloud Scoot API
func (h *Handler) SetSchedulerStatus(maxNumTasks int32) error {
	return schedthrift.SetSchedulerStatus(h.scheduler, maxNumTasks)
}

// GetClassLoadPcts Implements GetClassLoadPcts Cloud Scoot API
func (h *Handler) GetClassLoadPcts() (map[string]int32, error) {
	clp, err := schedthrift.GetClassLoadPcts(h.scheduler)
	log.Infof("GetClassLoadPcts returning: %v, err:%s", clp, err)
	return clp, err
}

// SetClassLoadPcts Implements SetClassLoadPcts Cloud Scoot API
func (h *Handler) SetClassLoadPcts(classLoadPcts map[string]int32) error {
	log.Infof("SetClassLoadPcts to %v", classLoadPcts)
	return schedthrift.SetClassLoadPcts(h.scheduler, classLoadPcts)
}

// GetRequestorToClassMap Implements GetRequestorToClassMap Cloud Scoot API
func (h *Handler) GetRequestorToClassMap() (map[string]string, error) {
	rm, err := schedthrift.GetRequestorToClassMap(h.scheduler)
	log.Infof("GetClassLoadPcts returning: %v, err:%s", rm, err)
	return rm, err
}

// SetRequestorToClassMap Implements SetRequestorToClassMap Cloud Scoot API
func (h *Handler) SetRequestorToClassMap(requestToClassMap map[string]string) error {
	log.Infof("SetRequestorToClassMap to %v", requestToClassMap)
	return schedthrift.SetRequestorToClassMap(h.scheduler, requestToClassMap)
}

// GetRebalanceMinDuration get the duration(minutes) that the scheduler needs to be exceeding
// the rebalance threshold before rebalancing.  <= 0 implies no rebalancing
func (h *Handler) GetRebalanceMinDuration() (int32, error) {
	d, err := schedthrift.GetRebalanceMinDuration(h.scheduler)
	log.Infof("GetRebalanceMinDuration returning: %d, err:%s", d, err)
	return d, err
}

// SetRebalanceMinDuration set the duration(minutes) that the scheduler needs to be exceeding
// the rebalance threshold before rebalancing.  <= 0 implies no rebalancing
func (h *Handler) SetRebalanceMinDuration(durationMin int32) error {
	log.Infof("SetRebalanceMinDuration to %d", durationMin)
	return schedthrift.SetRebalanceMinDuration(h.scheduler, durationMin)
}

// GetRebalanceThreshold the % spread threshold that must be exceeded to trigger rebalance
// <= 0 implies no rebalancing
func (h *Handler) GetRebalanceThreshold() (int32, error) {
	t, err := schedthrift.GetRebalanceThreshold(h.scheduler)
	log.Infof("GetRebalanceThreshold returning: %d, err:%s", t, err)
	return t, err
}

// SetRebalanceThreshold the % spread threshold that must be exceeded to trigger rebalance
// <= 0 implies no rebalancing
func (h *Handler) SetRebalanceThreshold(threshold int32) error {
	log.Infof("SetRebalanceThreshold to %d", threshold)
	return schedthrift.SetRebalanceThreshold(h.scheduler, threshold)
}
