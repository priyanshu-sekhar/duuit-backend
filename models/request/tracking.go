package request

import (
	"duuit/models/common"
	"time"
)

type TrackingRequest struct {
	GoalID     uint
	recordedAt time.Time
	SequenceNo int
	common.Status
}
