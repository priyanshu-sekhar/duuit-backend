package request

import (
	"duuit-backend/models/common"
	"time"
)

type TrackingRequest struct {
	GoalID     uint
	recordedAt time.Time
	SequenceNo int
	Status common.Status
}
