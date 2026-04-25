package codemode

import (
	"context"
	"sync/atomic"
	"time"
)

type HealthMonitor struct {
	client *Client
	ready  atomic.Bool
	stop   chan struct{}
}

func NewHealthMonitor(client *Client) *HealthMonitor {
	return &HealthMonitor{client: client, stop: make(chan struct{})}
}

func (h *HealthMonitor) Start(ctx context.Context, interval time.Duration) {
	go func() {
		t := time.NewTicker(interval)
		defer t.Stop()
		check := func() {
			cctx, cancel := context.WithTimeout(ctx, 2*time.Second)
			defer cancel()
			stats, err := h.client.HealthStats(cctx)
			if err != nil {
				h.ready.Store(false)
				return
			}
			h.ready.Store(true)
			JobsRunning.Set(float64(stats.JobsRunning))
			JobStoreSize.Set(float64(stats.JobStoreSize))
		}
		check()
		for {
			select {
			case <-ctx.Done():
				return
			case <-h.stop:
				return
			case <-t.C:
				check()
			}
		}
	}()
}

func (h *HealthMonitor) Stop() { close(h.stop) }

func (h *HealthMonitor) Ready() bool { return h.ready.Load() }
