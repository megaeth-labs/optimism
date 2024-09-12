package engine

import (
	"fmt"
	"time"

	"github.com/ethereum-optimism/optimism/op-service/eth"
)

type BuildStartedEvent struct {
	Info eth.PayloadInfo

	BuildStarted time.Time

	Parent eth.L2BlockRef

	// if payload should be promoted to safe (must also be pending safe, see DerivedFrom)
	IsLastInSpan bool
	// payload is promoted to pending-safe if non-zero
	DerivedFrom eth.L1BlockRef
}

func (ev BuildStartedEvent) String() string {
	return "build-started"
}

func (eq *EngDeriver) onBuildStarted(ev BuildStartedEvent) {
	// If a (pending) safe block, immediately seal the block
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.00000"), "onBuildStarted started: payloadId:", ev.Info.ID.String())
	if ev.DerivedFrom != (eth.L1BlockRef{}) {
		fmt.Println(time.Now().Format("2006-01-02 15:04:05.00000"), "emit BuildSealEvent: payloadId:", ev.Info.ID.String())
		eq.emitter.Emit(BuildSealEvent{
			Info:         ev.Info,
			BuildStarted: ev.BuildStarted,
			IsLastInSpan: ev.IsLastInSpan,
			DerivedFrom:  ev.DerivedFrom,
		})
	}
	fmt.Println(time.Now().Format("2006-01-02 15:04:05.00000"), "onBuildStarted ended: payloadId:", ev.Info.ID.String())
}
