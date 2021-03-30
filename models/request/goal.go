package request

type GoalRequest struct {
	UserID          uint
	Name            string
	TimelineInWeeks int
	Description     string
}
