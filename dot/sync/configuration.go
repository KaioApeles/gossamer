// Copyright 2024 ChainSafe Systems (ON)
// SPDX-License-Identifier: LGPL-3.0-only

package sync

import (
	"time"
)

type ServiceConfig func(svc *SyncService)

func WithStrategies(currentStrategy, defaultStrategy Strategy) ServiceConfig {
	return func(svc *SyncService) {
		svc.currentStrategy = currentStrategy
		svc.defaultStrategy = defaultStrategy

		wpCapacity := currentStrategy.NumOfTasks()
		if defaultStrategy != nil {
			wpCapacity = max(currentStrategy.NumOfTasks(), defaultStrategy.NumOfTasks())
		}
		wpCapacity *= 2 // add some buffer

		svc.workerPool = NewWorkerPool(WorkerPoolConfig{
			Capacity:   wpCapacity,
			MaxRetries: UnlimitedRetries,
		})
	}
}

func WithNetwork(net Network) ServiceConfig {
	return func(svc *SyncService) {
		svc.network = net
	}
}

func WithBlockState(bs BlockState) ServiceConfig {
	return func(svc *SyncService) {
		svc.blockState = bs
	}
}

func WithSlotDuration(slotDuration time.Duration) ServiceConfig {
	return func(svc *SyncService) {
		svc.slotDuration = slotDuration
	}
}

func WithMinPeers(min int) ServiceConfig {
	return func(svc *SyncService) {
		svc.minPeers = min
	}
}
