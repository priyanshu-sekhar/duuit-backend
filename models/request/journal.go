package request

import "duuit-backend/models/common"

type JournalRequest struct {
	TrackingID uint
	Privacy common.Privacy
	Text string
}
