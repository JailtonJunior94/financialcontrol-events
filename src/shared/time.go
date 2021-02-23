package shared

import "time"

type Time struct {
	time.Time
	Now      time.Time
	Location *time.Location
}

func NewTime(timeConfig ...Time) *Time {
	var newConfig Time

	if len(timeConfig) > 0 {
		newConfig = timeConfig[0]
		return &newConfig
	}

	return &Time{
		Now:      time.Now(),
		Location: time.Now().Location(),
	}
}

func (t *Time) DaysInMonth(i time.Time) int {
	return i.AddDate(0, 1, 0).Add(time.Nanosecond * -1).Day()
}

func (t *Time) StartDate() time.Time {
	return time.Date(t.Now.Year(), t.Now.Month(), 1, 0, 0, 0, 0, t.Location)
}

func (t *Time) EndDate() time.Time {
	return time.Date(t.Now.Year(), t.Now.Month(), t.DaysInMonth(t.StartDate()), 23, 59, 59, 59, t.Location)
}
