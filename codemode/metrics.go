package codemode

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	RunDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "codemode_run_duration_seconds",
		Help:    "Duration of code-mode tool calls.",
		Buckets: prometheus.DefBuckets,
	}, []string{"mode", "outcome"})

	JobsRunning = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "codemode_jobs_running",
		Help: "Number of running jobs in the workerd job store.",
	})

	JobStoreSize = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "codemode_jobstore_size",
		Help: "Total entries in the workerd job store.",
	})

	APICallDuration = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Name:    "codemode_api_call_duration_seconds",
		Help:    "Duration of api-proxy backend calls.",
		Buckets: prometheus.DefBuckets,
	}, []string{"status"})
)
