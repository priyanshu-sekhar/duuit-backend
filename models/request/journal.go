package request

import "duuit/models/common"

type JournalRequest struct {
	TrackingID uint
	Privacy common.Privacy
	Text string
}
